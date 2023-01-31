package main

import (
	"fmt"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/myfin"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/nacbank"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/tgbot"

	// "github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func main() {
	logger, err := log.New()
	if err != nil {
		fmt.Println("logger cannot be initialized")
	}

	cfg, err := config.New()
	if err != nil {
		logger.Error("config cannot be initialized")
	}

	app := fx.New(
		fx.Supply(logger, cfg),
		fx.Provide(
			tgbot.New,
			myfin.New,
			nacbank.New,
		),
	)

	fmt.Println("Победа")

	app.Run()

	// urlNBRB, telegram, err := config.New()
	// if err != nil {
	// 	logrus.Errorf("can't read config: %d", err)
	// }

	// chatID, err := strconv.Atoi(telegram.ChatID)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// tgClient, err := client.New(telegram.APIkey)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// rates, err := tgClient.GetRates(urlNBRB)
	// if err != nil {
	// 	logrus.Error(err)
	// }

	// var msg messages.Summary

	// summary, err := msg.GenerateFromRates(rates)
	// if err != nil {
	// 	logrus.Error(err)
	// }

	// Run cron schedule
	// crn := cron.New()
	// _, err = crn.AddFunc("@every 10s", func() {
	// 	_, err = tgClient.Bot.Send(tgAPI.NewMessage(int64(chatID), summary))
	// 	if err != nil {
	// 		logrus.Error(err)
	// 	}
	// })

	if err != nil {
		logrus.Error(err)
	}

	// for {
	// 	crn.Start()
	// }
}
