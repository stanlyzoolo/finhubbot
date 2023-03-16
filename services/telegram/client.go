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

type tgBot struct {
	bot *tgAPI.BotAPI
	log *log.Logger
	cfg *config.Config
}

func New(log *log.Logger, cfg *config.Config) (Bot, error) {
	bot, err := tgAPI.NewBotAPI(cfg.Telegram.APIkey)
	if err != nil {
		return nil, fmt.Errorf("can't initialize bot: %v", err)
	}

	keyboard := tgAPI.NewReplyKeyboard([]tgAPI.KeyboardButton{
		{
			Text: "Abrakadabra",
		},
		{
			Text: "Shvabra",
		},
	})

	keyboard.OneTimeKeyboard = true

	upd := tgAPI.NewUpdate(0)
	upd.Timeout = 60

	updates := bot.GetUpdatesChan(upd)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgAPI.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msg.ReplyMarkup = keyboard

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	return &tgBot{
		bot: bot,
		log: log,
		cfg: cfg,
	}, nil
}

func (b *tgBot) SendMessage(summary string) error {
	chatID, err := strconv.ParseInt(b.cfg.Telegram.ChatID, 10, 64)
	if err != nil {
		b.log.Errorf("can't parse chatID from env to int64: %v", err)

		return err
	}

	sentMsg, err := b.bot.Send(tgAPI.NewMessage(chatID, summary))
	b.log.Infof("Sent message: %s", sentMsg.Text)

	return err
}
