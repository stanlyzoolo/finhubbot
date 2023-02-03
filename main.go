package main

import (
	"fmt"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/myfin"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/nacbank"
	bot "github.com/stanlyzoolo/smartLaFamiliaBot/services/telegram"
	"github.com/stanlyzoolo/smartLaFamiliaBot/usecase"

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
		// fx.NopLogger,
		fx.Supply(logger, cfg),
		fx.Provide(
			bot.New,
			myfin.New,
			nacbank.New,
			usecase.New,
			services.New,
		),
		fx.Invoke(func(
			_ *usecase.Service,
		) {
		}),
	)

	app.Run()
}
