package config

import (
	"fmt"
	"os"
)

type (
	Config struct {
		MyFin    *MyFin
		NBRB     *NBRB
		Telegram *Telegram
	}

	MyFin struct {
		AllowedDomain string
		URL           string
	}

	NBRB struct {
		AllRatesURL string
		OneRateURL  string
	}

	Telegram struct {
		APIkey string
		ChatID string
	}
)

var ErrEmptyEnvVariable = fmt.Errorf("value is not set")

func New() (*Config, error) {
	// MyFin
	allowedDomain, ok := os.LookupEnv("ALLOWED_DOMAIN")
	if !ok || len(allowedDomain) == 0 {
		return nil, fmt.Errorf("ALLOWED_DOMAIN: %w", ErrEmptyEnvVariable)
	}

	url, ok := os.LookupEnv("MYFIN_URL")
	if !ok || len(url) == 0 {
		return nil, fmt.Errorf("MYFIN_URL: %w", ErrEmptyEnvVariable)
	}

	// NBRB
	allRatesURL, ok := os.LookupEnv("ALL_RATES_URL")
	if !ok || len(allRatesURL) == 0 {
		return nil, fmt.Errorf("ALL_RATES_URL: %w", ErrEmptyEnvVariable)
	}

	oneRateURL, ok := os.LookupEnv("ONE_RATE_URL")
	if !ok || len(oneRateURL) == 0 {
		return nil, fmt.Errorf("ONE_RATE_URL: %w", ErrEmptyEnvVariable)
	}

	// Telegram
	apiKey, ok := os.LookupEnv("TG_TOKEN")
	if !ok || len(apiKey) == 0 {
		return nil, fmt.Errorf("TG_TOKEN: %w", ErrEmptyEnvVariable)
	}

	chatID, ok := os.LookupEnv("CHAT_ID")
	if !ok || len(chatID) == 0 {
		return nil, fmt.Errorf("CHAT_ID: %w", ErrEmptyEnvVariable)
	}

	return &Config{
			MyFin: &MyFin{
				AllowedDomain: allowedDomain,
				URL:           url,
			},
			NBRB: &NBRB{
				AllRatesURL: allRatesURL,
				OneRateURL:  oneRateURL,
			},
			Telegram: &Telegram{
				APIkey: apiKey,
				ChatID: chatID,
			},
		},
		nil
}
