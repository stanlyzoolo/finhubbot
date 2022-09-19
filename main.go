package main

import (
	"encoding/json"
	"fmt"
	"io"

	// TODO Использовать логгер: https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api/v5#SetLogger

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logrus.Errorf("can't read response: %d", err)
	}

	curReq, err := currencies.CurrencyReq(cfg.OneRateURL, config.USD)
	if err != nil {
		fmt.Printf("Bad news: %s", err.Error())
	}

	// TODO вынести токен в переменную окружения
	bot, err := tgbotapi.NewBotAPI("5677404105:AAEGlBwarltHXGSzvjJOxZbXNGTatij_98w")
	if err != nil {
		panic(err)
	}

	resp, err := bot.Client.Do(curReq)
	if err != nil {
		logrus.Errorf("can't Do request: %d", err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("can't read response: %d", err)
	}

	var cur currencies.Currency
	if err := json.Unmarshal(respBody, &cur); err != nil {
		logrus.Errorf("can't Unmarshal body: %d", err)
	}

	// Instead of typing the API token directly into the file, we're using environment variables.
	// This makes it easy to configure our Bot to use the right account and prevents us from
	// leaking our real token into the world.
	// Anyone with your token can send and receive messages from your Bot!
	// my token: 5677404105:AAEGlBwarltHXGSzvjJOxZbXNGTatij_98w

	bot.Debug = true
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		rates := fmt.Sprint(cur.Name, cur.OfficialRate)

		// Now that we know we've gotten a new message, we can construct a
		// reply! We'll take the Chat ID and Text from the incoming message
		// and use it to create a new message.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, rates)
		// We'll also say that this message is a reply to the previous message.
		// For any other specifications than Chat ID or Text, you'll need to
		// set fields on the `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}
}
