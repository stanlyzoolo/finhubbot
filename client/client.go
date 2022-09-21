package client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/juju/errors"
)

type Client struct {
	Bot *tgbotapi.BotAPI
}

func New(token string) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errors.Errorf("can't initialize bot: %d", err)
	}

	return &Client{
		Bot: bot,
	}, nil
}
