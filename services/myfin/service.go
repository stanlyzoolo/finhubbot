package myfin

import (
	"database/sql"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/repo"

	"github.com/gocolly/colly"
)

type Service interface {
	scrapDomain() ([]string, error)
	storeRates([]string) error
}

type service struct {
	c       *colly.Collector
	log     *log.Logger
	cfg     *config.Config
	storage repo.Commercials
}

func NewService(log *log.Logger, cfg *config.Config, db *sql.DB) Service {
	srv := &service{
		c:       colly.NewCollector(),
		log:     log,
		cfg:     cfg,
		storage: repo.NewCommercials(db, log),
	}

	go func() {
		srv.runBySchedule()
	}()

	return srv
}

type (
	Currency struct {
		Bank Bank
		USD  USD
		EUR  EUR
		RUB  RUB
	}

	Bank struct {
		Name string
	}
	USD struct {
		Buying, Selling string
	}
	EUR struct {
		Buying, Selling string
	}
	RUB struct {
		Buying, Selling string
	}
)
