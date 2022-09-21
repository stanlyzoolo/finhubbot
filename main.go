package main

import (
	"fmt"
	"strconv"

	// TODO Использовать логгер: https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api/v5#SetLogger

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/stanlyzoolo/smartLaFamiliaBot/client"
	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/messages"
)

func main() {
	urlNBRB, telegram, err := config.New()
	if err != nil {
		logrus.Errorf("can't read config: %d", err)
	}

	chatID, err := strconv.Atoi(telegram.ChatID)
	if err != nil {
		fmt.Println(err)
	}

	// Init bot
	tgClient, err := client.New(telegram.APIkey)
	if err != nil {
		fmt.Println(err)
	}

	rates, err := tgClient.GetRates(urlNBRB) // panic
	if err != nil {
		logrus.Error(err)
	}

	var msg messages.Summary

	summary, err := msg.GenerateFromRates(rates)
	if err != nil {
		logrus.Error(err)
	}

	// Run cron schedule
	crn := cron.New()
	_, err = crn.AddFunc("@every 30s", func() {
		_, err = tgClient.Bot.Send(tgbotapi.NewMessage(int64(chatID), summary))
		if err != nil {
			logrus.Error(err)
		}
	})

	if err != nil {
		logrus.Error(err)
	}

	for {
		crn.Start()
	}
}
