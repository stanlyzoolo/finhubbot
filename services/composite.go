package services

import (
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/myfin"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/nacbank"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/telegram"
)

type Composite interface {
	MyFin() myfin.Service
	NatBank() nacbank.NBRB
	TelBot() telegram.Bot
}

type composite struct {
	Myfin    myfin.Service
	NBRB     nacbank.NBRB
	Telegram telegram.Bot
}

// MyFin implements Composite
func (c *composite) MyFin() myfin.Service {
	return c.Myfin
}

// NatBank implements Composite
func (c *composite) NatBank() nacbank.NBRB {
	return c.NBRB
}

// TelBot implements Composite
func (c *composite) TelBot() telegram.Bot {
	return c.Telegram
}

func New(useMyfin myfin.Service, useNBRB nacbank.NBRB, useBot telegram.Bot) Composite {
	return &composite{
		Myfin:    useMyfin,
		NBRB:     useNBRB,
		Telegram: useBot,
	}
}
