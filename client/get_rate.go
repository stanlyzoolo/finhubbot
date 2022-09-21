package client

import (
	"github.com/sirupsen/logrus"
	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/currencies"
)

func (c *Client) GetRates(cfg *config.NBRB) ([]currencies.Rate, error) {
	rates := make([]currencies.Rate, 0)

	for code, flag := range currencies.CodesAndFlags {
		req, err := getCurrency(cfg.OneRateURL, code)
		if err != nil {
			logrus.Errorf("bad news: %d", err)
		}

		resp, err := c.Bot.Client.Do(req)
		if err != nil {
			logrus.Errorf("can't Do request: %d", err)
		}

		rate, err := readResponse(resp)
		if err != nil {
			logrus.Errorf("can't read response: %d", err)
		}

		rate.CountryFlagUTF8 = flag

		rates = append(rates, *rate)
	}

	return rates, nil
}
