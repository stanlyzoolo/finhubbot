package usecase

import (
	"context"

	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/messages"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services"

	"github.com/robfig/cron/v3"
)

type Service interface {
	GenerateNatBankRatesInfo() (string, error)
	RunByCron() error
}

type service struct {
	log     *log.Logger
	clients services.Composite
}

func New(log *log.Logger, clients services.Composite) Service {
	return &service{
		log:     log,
		clients: clients,
	}
}

func (srv *service) GenerateNatBankRatesInfo() (string, error) {
	rates, err := srv.clients.NatBank().GetRates(context.Background())
	if err != nil {
		srv.log.Errorf("can't rates from National Bank: %w", err)

		return "", nil
	}

	ready, err := messages.GenerateFromNatBankRates(rates)
	if err != nil {
		srv.log.Errorf("can't construct summary from rates: %w", err)

		return "", nil
	}

	return ready, err
}

// TODO Доделать и запустить
func (srv *service) RunByCron() error {
	summary, err := srv.GenerateNatBankRatesInfo()
	if err != nil {
		srv.log.Error(err)

		return err
	}

	// Run cron schedule
	crn := cron.New()
	_, err = crn.AddFunc("@every 10s", func() {
		err = srv.clients.TelBot().SendMessage(summary)
		if err != nil {
			srv.log.Error(err)
		}
	})

	for {
		crn.Start()
	}

	return nil
}
