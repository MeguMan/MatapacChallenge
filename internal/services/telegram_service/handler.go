package telegram_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/MeguMan/MatapacChallenge/internal/storage"
	"github.com/gagliardetto/solana-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
			//case update.Message.Command() == "update":
			//	go s.update(update)
			case update.Message.Command() == "top":
				go s.top(update)
			case update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text == addText:
				go s.addUser(ctx, update)
			//case update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text == updateText:
			//	go s.updateUser(ctx, update)
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

func (s *service) update(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, updateText)
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
		s.sendMsg(update, invalidPublicKeyErrText, "")
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
			s.sendMsg(update, userAlreadyExistsErrText, "")
		} else {
			s.sendMsg(update, internalErrorText, "")
		}
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, addSuccessText)
	if _, err := s.bot.Send(msg); err != nil {
		fmt.Println(err)
		return
	}

	msg = tgbotapi.NewMessage(487861234, fmt.Sprintf("(%d,'%s','%s')", update.Message.From.ID, update.Message.From.UserName, publicKey.String()))
	if _, err := s.bot.Send(msg); err != nil {
		fmt.Println(err)
		return
	}
}

func (s *service) updateUser(ctx context.Context, update tgbotapi.Update) {
	publicKey, err := solana.PublicKeyFromBase58(update.Message.Text)
	if err != nil {
		s.sendMsg(update, invalidPublicKeyErrText, "")
		return
	}

	err = s.storageService.UpdateUser(ctx, storage.User{
		TgID:         update.Message.From.ID,
		SolPublicKey: publicKey.String(),
	})
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, storage.ErrUniqueKeyViolation) {
			s.sendMsg(update, userAlreadyExistsErrText, "")
		} else {
			s.sendMsg(update, internalErrorText, "")
		}
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, updateSuccessText)
	if _, err := s.bot.Send(msg); err != nil {
		fmt.Println(err)
		return
	}

	msg = tgbotapi.NewMessage(487861234, fmt.Sprintf("%s changed public key to %s", update.Message.From.UserName, publicKey.String()))
	if _, err := s.bot.Send(msg); err != nil {
		fmt.Println(err)
		return
	}
}

func (s *service) top(update tgbotapi.Update) {
	s.sendMsg(update, s.getTopMsg(), tgbotapi.ModeHTML)
}

func (s *service) sendMsg(update tgbotapi.Update, msgText string, parseMode string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)

	if parseMode != "" {
		msg.ParseMode = parseMode
	}

	if _, err := s.bot.Send(msg); err != nil {
		fmt.Println(err)
		return
	}
}
