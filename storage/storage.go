package storage

import (
	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/repo"

	"go.uber.org/fx"
)

func Construct() fx.Option {
	return fx.Provide(
		repo.NewCommercials,
		repo.NewNBRB,
	)
}

type Storage interface {
	Commercials() repo.Commercials
	NatBankRB() repo.NatBankRB
}

type storage struct {
	repoCommercials repo.Commercials
	repoNatBankRB   repo.NatBankRB
}

func New(c repo.Commercials, n repo.NatBankRB) Storage {
	return &storage{
		repoCommercials: c,
		repoNatBankRB:   n,
	}
}

func (s *storage) Commercials() repo.Commercials {
	return s.repoCommercials
}

func (s *storage) NatBankRB() repo.NatBankRB {
	return s.repoNatBankRB
}
