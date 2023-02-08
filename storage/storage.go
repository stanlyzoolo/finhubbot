package storage

import (
	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/repo"
	"go.uber.org/fx"
)

func Construct() fx.Option {
	return fx.Provide(
		repo.NewCommercials,
	)
}

type Storage interface {
	Commercials() repo.Commercials
}

type storage struct {
	repoCommercials repo.Commercials
}

func New(c repo.Commercials) Storage {
	return &storage{repoCommercials: c}
}

func (c *storage) Commercials() repo.Commercials {
	return c.repoCommercials
}
