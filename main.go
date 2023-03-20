package main

import (
	"database/sql"
	"fmt"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/schema"

	"github.com/stanlyzoolo/smartLaFamiliaBot/services/myfin"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/natbank"
	bot "github.com/stanlyzoolo/smartLaFamiliaBot/services/telegram"
	"github.com/stanlyzoolo/smartLaFamiliaBot/storage"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

func main() {
	logger, err := log.New()
	if err != nil {
		fmt.Println("logger cannot be initialized")
	}

	cfg, err := config.New()
	if err != nil {
		logger.Errorf("config cannot be initialized: %v", err)
	}

	db, err := sql.Open("sqlite3", "./schema/rates.db")
	if err != nil {
		logger.Error("db: cannot open connection")
	}

	if err = db.Ping(); err != nil {
		logger.Error(err)
	}

	_ = schema.New(logger, db)

	app := fx.New(
		// fx.NopLogger,
		fx.Supply(logger, cfg, db),
		fx.Provide(
			bot.New,
			storage.New,
			natbank.NewService,
			myfin.NewService,
		),
		storage.Construct(),
		// TODO Не работает этот вызов, скорее всего из за бесконечного цикла в одном из сервисов
		fx.Invoke(func(
			_ natbank.Service,
			_ myfin.Service,
		) {
		}),
	)

	app.Run()
}
