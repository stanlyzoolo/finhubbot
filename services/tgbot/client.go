package tgbot

import (
	"fmt"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"

	tgAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	Bot    *tgAPI.BotAPI
	Logger *log.Logger
	cfg    *config.Config
}

func New(log *log.Logger, cfg *config.Config) (*Client, error) {
	bot, err := tgAPI.NewBotAPI(cfg.Telegram.APIkey)
	if err != nil {
		return nil, fmt.Errorf("can't initialize bot: %w", err)
	}

	return &Client{
		Bot:    bot,
		Logger: log,
		cfg:    cfg,
	}, nil
}
