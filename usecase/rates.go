package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/messages"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services"

	"go.uber.org/fx"
)

// type Service interface {
// 	GenerateNatBankRatesInfo(ctx context.Context) (string, error)
// 	RunByCron(ctx context.Context) error
// }

type Service struct {
	log     *log.Logger
	clients services.Composite
}

func New(lc fx.Lifecycle, log *log.Logger, clients services.Composite) *Service {
	srv := &Service{
		log:     log,
		clients: clients,
	}

	srvCtx := context.Background()

	lc.Append(fx.Hook{
		OnStart: func(srvCtx context.Context) error {
			go func() {
				err := srv.RunByCron(srvCtx)
				srv.log.Error(err)
			}()

			return nil
		},

		OnStop: func(_ context.Context) error {
			srvCtx.Done()

			return nil
		},
	})

	return srv
}

func (srv *Service) RunByCron(ctx context.Context) error {
	summary, err := srv.GenerateNatBankRatesInfo(ctx)
	if err != nil {
		srv.log.Error(err)

		return err
	}

	t := time.NewTicker(time.Second * 10)

	for range t.C {
		err = srv.clients.TelBot().SendMessage(summary)
		if err != nil {
			srv.log.Error(err)

			return err
		}

		// TODO инфа приходит - написать комбайн для сообщения в телегу
		_, err = srv.GenerateCommercialBanksInfo(ctx)
	}

	return err
}

func (srv *Service) GenerateNatBankRatesInfo(ctx context.Context) (string, error) {
	rates, err := srv.clients.NatBank().GetRates(ctx)
	if err != nil {
		srv.log.Errorf("can't rates from National Bank: %v", err)

		return "", nil
	}

	ready, err := messages.GenerateSummaryFromNatBank(rates)
	if err != nil {
		srv.log.Errorf("can't construct summary from rates: %v", err)

		return "", nil
	}

	return ready, err
}

func (srv *Service) GenerateCommercialBanksInfo(ctx context.Context) (string, error) {
	if err := srv.clients.MyFin().SetAllowedDomain(); err != nil {
		srv.log.Errorf("can't set allowed domain: %v", err)

		return "", fmt.Errorf("can't set allowed domain: %v", err)
	}

	commercilaRates, err := srv.clients.MyFin().ScrapDomain()
	if err != nil {
		srv.log.Errorf("can't scrap commercial rates from established domain: %v", err)

		return "", fmt.Errorf("can't scrap commercial rates from established domain: %v", err)
	}

	ordered := srv.clients.MyFin().OrderIncomingData(commercilaRates)

	srv.log.Infof("commercial rates:\n %v", ordered)

	return "", nil
}
