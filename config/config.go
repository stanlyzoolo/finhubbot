package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Config struct {
		MyFin    *MyFin
		NBRB     *NBRB
		Telegram *Telegram
	}

	MyFin struct {
		Delay           int
		Every           int
		IterationsCount int
		URL             string
		AllowedDomain   string
		AllowedWeekdays []string
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

func New() (*Config, error) { // nolint
	// MyFin
	allowedDomain, ok := os.LookupEnv("ALLOWED_DOMAIN")
	if !ok || len(allowedDomain) == 0 {
		return nil, fmt.Errorf("ALLOWED_DOMAIN: %w", ErrEmptyEnvVariable)
	}

	url, ok := os.LookupEnv("MYFIN_URL")
	if !ok || len(url) == 0 {
		return nil, fmt.Errorf("MYFIN_URL: %w", ErrEmptyEnvVariable)
	}

	count, ok := os.LookupEnv("ITERATIONS_COUNT")
	if !ok || len(count) == 0 {
		return nil, fmt.Errorf("ITERATIONS_COUNT: %w", ErrEmptyEnvVariable)
	}

	icount, err := strconv.Atoi(count)
	if err != nil {
		return nil, fmt.Errorf("failed to convert COUNT to int type: %w", err)
	}

	delay, ok := os.LookupEnv("DELAY")
	if !ok || len(delay) == 0 {
		return nil, fmt.Errorf("DELAY: %w", ErrEmptyEnvVariable)
	}

	idelay, err := strconv.Atoi(delay)
	if err != nil {
		return nil, fmt.Errorf("failed to convert DELAY to int type: %w", err)
	}

	every, ok := os.LookupEnv("EVERY")
	if !ok || len(every) == 0 {
		return nil, fmt.Errorf("EVERY: %w", ErrEmptyEnvVariable)
	}

	ievery, err := strconv.Atoi(every)
	if err != nil {
		return nil, fmt.Errorf("failed to convert EVERY to int type: %w", err)
	}

	allowedWeekdays, ok := os.LookupEnv("ALLOWED_WEEKDAYS")
	if !ok || len(allowedWeekdays) == 0 {
		return nil, fmt.Errorf("ALLOWED_WEEKDAYS: %w", ErrEmptyEnvVariable)
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
				Delay:           idelay,
				Every:           ievery,
				IterationsCount: icount,
				URL:             url,
				AllowedDomain:   allowedDomain,
				AllowedWeekdays: strings.Split(allowedWeekdays, " "),
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
		err
}
