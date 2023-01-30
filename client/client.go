package client

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	Bot *tgbotapi.BotAPI
}

func New(token string) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("can't initialize bot: %w", err)
	}

	return &Client{
		Bot: bot,
	}, nil
}
