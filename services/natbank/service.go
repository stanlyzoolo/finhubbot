package natbank

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/myfin"
	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/repo"
)

// National Bank Republic of Belarus
type Service interface {
	getCurrenciesRates(ctx context.Context) ([]Rate, error)
	storeRates(ctx context.Context, rates []Rate) error
}

type service struct {
	client  *http.Client
	log     *log.Logger
	cfg     *config.Config
	storage repo.NatBankRB
}

func NewService(log *log.Logger, cfg *config.Config, db *sql.DB) Service {
	srv := &service{
		client:  &http.Client{},
		log:     log,
		cfg:     cfg,
		storage: repo.NewNBRB(db, log),
	}

	delay := time.Second * time.Duration(srv.cfg.MyFin.Delay)

	go func() {
		srv.run(delay)
	}()

	return srv
}

func (s *service) run(delay time.Duration) {
	t := time.NewTicker(time.Second)

	ctx := context.Background()

	for {
		select {
		case <-ctx.Done():
			s.log.Info("Done")

			return
		case <-t.C:
			time.Sleep(delay)

			today := time.Now().Weekday().String()

			if s.allowedWeekdays().Has(today) {
				for i := 1; i <= len(codesAndFlags); i++ {
					rates, err := s.getCurrenciesRates(ctx)
					if err != nil {
						s.log.Error(err)
					}

					if err = s.storeRates(ctx, rates); err != nil {
						s.log.Error(err)

						return
					}
				}
			}
		}
	}
}

func (s *service) allowedWeekdays() *myfin.Weekdays {
	weekdays := make(myfin.Weekdays, len(s.cfg.MyFin.AllowedWeekdays))

	for i, day := range s.cfg.MyFin.AllowedWeekdays {
		weekdays[day] = myfin.Weekday(i)
	}

	return &weekdays
}
