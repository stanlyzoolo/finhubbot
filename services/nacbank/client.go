package nacbank

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stanlyzoolo/smartLaFamiliaBot/banks/nacbank"
	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
)

// National Bank Republic of Belarus
type NBRB interface {
	GetRates(ctx context.Context) ([]nacbank.Rate, error)
}

type nbrb struct {
	client *http.Client
	log    *log.Logger
	cfg    *config.Config
}

func New(log *log.Logger, cfg *config.Config) NBRB {
	return &nbrb{
		client: http.DefaultClient,
		log:    log,
		cfg:    cfg,
	}
}

// TODO refactor
func (n *nbrb) GetRates(ctx context.Context) ([]nacbank.Rate, error) {
	rates := make([]nacbank.Rate, 0)

	for code, flag := range nacbank.CodesAndFlags {
		req, err := preparetRequest(ctx, n.cfg.NBRB.OneRateURL, code)
		if err != nil {
			n.log.Error(err)

			return nil, fmt.Errorf("%v", err)
		}

		resp, err := n.client.Do(req)
		if err != nil {
			n.log.Errorf("can't Do request: %v", err)

			return nil, fmt.Errorf("can't Do request: %v", err)
		}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			n.log.Errorf("can't read response: %v", err)

			return nil, fmt.Errorf("can't read response: %v", err)
		}

		defer resp.Body.Close()

		var rate nacbank.Rate

		if err := json.Unmarshal(respBody, &rate); err != nil {
			return nil, fmt.Errorf("can't unmarshal body: %v", err)
		}

		rate.CountryFlagUTF8 = flag

		rates = append(rates, rate)
	}

	return rates, nil
}

// Формат запроса в НБ РБ:
// https://www.nbrb.by/api/exrates/rates/840?parammode=1
func preparetRequest(ctx context.Context, url string, curID int) (*http.Request, error) {
	var body io.Reader

	urlWithCurrency := fmt.Sprint(url, curID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlWithCurrency, body)
	if err != nil {
		return nil, fmt.Errorf("can't set request: %v", err)
	}

	q := req.URL.Query()
	q.Add("parammode", "1")

	req.URL.RawQuery = q.Encode()

	return req, nil
}
