package telegram

import (
	"fmt"
	"strconv"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"

	tgAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot interface {
	SendMessage(summary string) error
}

type bot struct {
	bot *tgAPI.BotAPI
	log *log.Logger
	cfg *config.Config
}

func New(log *log.Logger, cfg *config.Config) (Bot, error) {
	botAPI, err := tgAPI.NewBotAPI(cfg.Telegram.APIkey)
	if err != nil {
		return nil, fmt.Errorf("can't initialize bot: %w", err)
	}

	return &bot{
		bot: botAPI,
		log: log,
		cfg: cfg,
	}, nil
}

func (b *bot) SendMessage(summary string) error {
	chatID, err := strconv.ParseInt(b.cfg.Telegram.ChatID, 10, 64)
	if err != nil {
		b.log.Errorf("can't parse chatID from env to int64: %w", err)

		return err
	}

	sentMsg, err := b.bot.Send(tgAPI.NewMessage(chatID, summary))
	b.log.Infof("Sent message: %s", sentMsg.Text)

	return err
}
