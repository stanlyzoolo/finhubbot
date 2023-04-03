package telegram

import (
	"fmt"
	"strconv"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot interface {
	SendMessage(summary string) error
}

type tgBot struct {
	bot *tgbotapi.BotAPI
	log *log.Logger
	cfg *config.Config
}

func New(log *log.Logger, cfg *config.Config) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.APIkey)
	if err != nil {
		return nil, fmt.Errorf("can't initialize bot: %w", err)
	}

	// []tgAPI.KeyboardButton{
	// 	{
	// 		Text: "USD",
	// 	},
	// 	{
	// 		Text: "Евро",
	// 	},
	// 	{
	// 		Text: "Росс. рубль",
	// 	},
	// },
	// keyboard := tgAPI.NewReplyKeyboard([]tgAPI.KeyboardButton{
	// 	{
	// 		Text: "Abrakadabra",
	// 	},
	// 	{
	// 		Text: "Shvabra",
	// 	},
	// })

	// keyboard.OneTimeKeyboard = true

	var keyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Доллар США"),
			tgbotapi.NewKeyboardButton("Евро"),
			tgbotapi.NewKeyboardButton("Руб"),
		),
	)

	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60

	updates := bot.GetUpdatesChan(upd)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

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

	sentMsg, err := b.bot.Send(tgbotapi.NewMessage(chatID, summary))
	b.log.Infof("Sent message: %s", sentMsg.Text)

	return err
}

var mainKeyboard string

var backButton string

// TODO преобразовать в структуру с вложенными клавиатурами по слоям
var natBankLayerKeyboard string

// TODO преобразовать в структуру с вложенными клавиатурами по слоям
var commercialsLayerKeyboard string
