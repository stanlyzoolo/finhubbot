package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/juju/errors"
)

// API НацБанка РБ
type NBRB struct {
	AllRatesURL string
	OneRateURL  string
}

type BotToken struct {
	Token string
}

func New() (*NBRB, *BotToken, error) {
	if err := godotenv.Load(); err != nil {
		return nil, nil, errors.New("can't load .env file")
	}

	return &NBRB{
			AllRatesURL: os.Getenv("AllRatesURL"),
			OneRateURL:  os.Getenv("OneRateURL"),
		},
		&BotToken{
			Token: os.Getenv("TGbotTOKEN"),
		},
		nil
}
