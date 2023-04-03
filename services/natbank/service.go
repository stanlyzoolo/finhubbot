package natbank

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/robfig/cron/v3"
	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/repo"
)

// National Bank Republic of Belarus
type Service interface {
	getCurrenciesRates(ctx context.Context) ([]Rate, error)
	storeRates(ctx context.Context, rates []Rate) error
}

type service struct {
	client  *http.Client
	cr      *cron.Cron
	log     *log.Logger
	cfg     *config.Config
	storage repo.NatBankRB
}

func NewService(log *log.Logger, cfg *config.Config, db *sql.DB) Service {
	cr := cron.New()

	srv := &service{
		client:  &http.Client{},
		cr:      cr,
		log:     log,
		cfg:     cfg,
		storage: repo.NewNBRB(db, log),
	}

	run := func() {
		srv.run()
	}

	if _, err := cr.AddFunc("10 11 * * MON-FRI", run); err != nil {
		srv.log.Error(err)
	}

	go func() {
		cr.Run()
	}()

	return srv
}

func (s *service) run() {
	rates, err := s.getCurrenciesRates(context.Background())
	if err != nil {
		s.log.Error(err)
	}

	if err = s.storeRates(context.Background(), rates); err != nil {
		s.log.Error(err)

		return
	}
}
