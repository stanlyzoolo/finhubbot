package config

import (
	"github.com/joho/godotenv"
	"github.com/juju/errors"
	// "golang.org/x/net/html"

	"os"
)

// API НацБанка РБ
type NBRB struct {
	AllRatesURL string
	OneRateURL  string
}

type Telegram struct {
	APIkey string
	ChatID string
}

func New() (*NBRB, *Telegram, error) {
	if err := godotenv.Load(); err != nil {
		return nil, nil, errors.New("can't load .env file")
	}

	return &NBRB{
			AllRatesURL: os.Getenv("AllRatesURL"),
			OneRateURL:  os.Getenv("OneRateURL"),
		},
		&Telegram{
			APIkey: os.Getenv("TGbotTOKEN"),
			ChatID: os.Getenv("ChatID"),
		},
		nil
}

func htmlParse() {
	// html.
}
