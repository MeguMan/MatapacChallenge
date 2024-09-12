package telegram_service

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Service interface {
	Handle(ctx context.Context) error
}

type service struct {
	chainstackService chainstackService
	storageService    storageService
	updates           tgbotapi.UpdatesChannel
	bot               *tgbotapi.BotAPI
}

func New(chainstackService chainstackService, storageService storageService, telegramBotToken string) (Service, error) {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		return nil, err
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}

	return &service{
		chainstackService: chainstackService,
		storageService:    storageService,
		updates:           updates,
		bot:               bot,
	}, nil
}
