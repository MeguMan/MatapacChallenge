package telegram_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/MeguMan/MatapacChallenge/internal/storage"
	"github.com/gagliardetto/solana-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math"
	"os"
	"sort"
)

func (s *service) Handle(ctx context.Context) error {
	for update := range s.updates {
		update := update

		if update.Message != nil {
			switch {
			case update.Message.Command() == "start":
				go s.start(update)
			case update.Message.Command() == "add":
				go s.add(update)
			case update.Message.Command() == "top":
				go s.top(ctx, update)
			case update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text == addText:
				go s.addUser(ctx, update)
			default:
				continue
			}
		}
	}

	return nil
}

func (s *service) start(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, startText)
	if _, err := s.bot.Send(msg); err != nil {
		fmt.Println(err)
		return
	}
}

func (s *service) add(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, addText)
	msg.ReplyMarkup = tgbotapi.ForceReply{
		ForceReply: true,
	}
	if _, err := s.bot.Send(msg); err != nil {
		fmt.Println(err)
		return
	}
}

func (s *service) addUser(ctx context.Context, update tgbotapi.Update) {
	publicKey, err := solana.PublicKeyFromBase58(update.Message.Text)
	if err != nil {
		s.sendMsg(update, invalidPublicKeyErrText)
		return
	}

	err = s.storageService.CreateUser(ctx, storage.User{
		TgID:         update.Message.From.ID,
		TgUsername:   update.Message.From.UserName,
		SolPublicKey: publicKey.String(),
	})
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, storage.ErrUniqueKeyViolation) {
			s.sendMsg(update, userAlreadyExistsErrText)
		} else {
			s.sendMsg(update, internalErrorText)
		}
		return
	}

	f, err := os.OpenFile("./../../../users.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Fprintln(f, fmt.Sprintf("%d-%s-%s", update.Message.From.ID, update.Message.From.UserName, publicKey.String()))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, addSuccessText)
	if _, err := s.bot.Send(msg); err != nil {
		fmt.Println(err)
		return
	}
}

func (s *service) top(ctx context.Context, update tgbotapi.Update) {
	users, err := s.storageService.GetUsersSolAccounts(ctx)
	if err != nil {
		fmt.Println(err)
		s.sendMsg(update, internalErrorText)
		return
	}

	mpUserNameByPublicKey := make(map[string]string, len(users))
	publicKeys := make([]string, 0, len(users))
	for _, user := range users {
		mpUserNameByPublicKey[user.SolPublicKey] = user.TgUsername
		publicKeys = append(publicKeys, user.SolPublicKey)
	}

	accounts, err := s.chainstackService.GetAccountsBalance(ctx, publicKeys)
	if err != nil {
		fmt.Println(err)
		s.sendMsg(update, internalErrorText)
		return
	}

	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].Sol > accounts[j].Sol
	})

	textMsg := ""
	for i, account := range accounts {
		textMsg += fmt.Sprintf("%d. %s - %f\n", i+1, mpUserNameByPublicKey[account.PublicKey], math.Round(account.Sol*100)/100)
	}

	s.sendMsg(update, textMsg)
}

func (s *service) sendMsg(update tgbotapi.Update, msgText string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err := s.bot.Send(msg); err != nil {
		fmt.Println(err)
		return
	}
}
