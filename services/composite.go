package services

import (
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/myfin"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/nacbank"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/telegram"
)

type Composite interface {
	MyFin() myfin.Myfin
	NatBank() nacbank.NBRB
	TelBot() telegram.Bot
}

type composite struct {
	Myfin    myfin.Myfin
	NBRB     nacbank.NBRB
	Telegram telegram.Bot
}

// MyFin implements Composite
func (c *composite) MyFin() myfin.Myfin {
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

func New(useMyfin myfin.Myfin, useNBRB nacbank.NBRB, useBot telegram.Bot) Composite {
	return &composite{
		Myfin:    useMyfin,
		NBRB:     useNBRB,
		Telegram: useBot,
	}
}
