package client

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/stanlyzoolo/smartLaFamiliaBot/banks/nacbank"
	"github.com/stanlyzoolo/smartLaFamiliaBot/config"

	// TODO Переделать на fmt пакет
	"github.com/juju/errors"
)

func (c *Client) GetRates(cfg *config.NBRB) ([]nacbank.Rate, error) {
	rates := make([]nacbank.Rate, 0)

	client := http.Client{}

	for code, flag := range nacbank.CodesAndFlags {
		req, err := getRequest(cfg.OneRateURL, code)
		if err != nil {
			return nil, errors.Errorf("bad news: %d", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.Errorf("can't Do request: %s", err.Error())
		}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Errorf("can't read response: %d", err)
		}

		defer resp.Body.Close()

		var rate nacbank.Rate

		if err := json.Unmarshal(respBody, &rate); err != nil {
			return nil, errors.Errorf("can't unmarshal body: %d", err)
		}

		rate.CountryFlagUTF8 = flag

		rates = append(rates, rate)
	}

	return rates, nil
}
