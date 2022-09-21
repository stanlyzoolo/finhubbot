package client

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/stanlyzoolo/smartLaFamiliaBot/currencies"
)

func readResponse(resp *http.Response) (*currencies.Rate, error) {
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("can't read response: %d", err)
	}

	var cur currencies.Currency

	if err := json.Unmarshal(respBody, &cur); err != nil {
		logrus.Errorf("can't unmarshal body: %d", err)
	}

	return &currencies.Rate{
		Currency: cur,
	}, nil
}
