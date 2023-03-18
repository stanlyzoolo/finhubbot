package nacbank

import (
	"context"
	"net/http"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/repo"
)

// National Bank Republic of Belarus
type NBRB interface {
	GetCurrencyRate(ctx context.Context) ([]Rate, error)
	StoreRates(ctx context.Context, rates []Rate) error
}

type client struct {
	client  *http.Client
	log     *log.Logger
	cfg     *config.Config
	storage repo.NatBankRB
}

func New(log *log.Logger, cfg *config.Config) NBRB {
	return &client{
		client: http.DefaultClient,
		log:    log,
		cfg:    cfg,
	}
}
