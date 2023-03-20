package services

import (
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/myfin"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/natbank"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/telegram"
)

type Composite interface {
	MyFin() myfin.Service
	NatBank() natbank.Service
	TelBot() telegram.Bot
}

type composite struct {
	Myfin    myfin.Service
	NBRB     natbank.Service
	Telegram telegram.Bot
}

// MyFin implements Composite
func (c *composite) MyFin() myfin.Service {
	return c.Myfin
}

// NatBank implements Composite
func (c *composite) NatBank() natbank.Service {
	return c.NBRB
}

// TelBot implements Composite
func (c *composite) TelBot() telegram.Bot {
	return c.Telegram
}

func New(myfin myfin.Service, natbank natbank.Service, bot telegram.Bot) Composite {
	c := &composite{
		Myfin:    myfin,
		NBRB:     natbank,
		Telegram: bot,
	}

	return c
}
