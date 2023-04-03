package repo

import (
	"context"
	"database/sql"

	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/dbo"

	_ "github.com/mattn/go-sqlite3" // nolint
)

type Commercials interface {
	Create(ctx context.Context, dbo dbo.CommercialRate) error
}

type commercials struct {
	db  *sql.DB
	log *log.Logger
}

func NewCommercials(db *sql.DB, log *log.Logger) Commercials {
	return &commercials{
		db:  db,
		log: log,
	}
}

// TODO Может всё таки возвращать ID --> returning id
func (c *commercials) Create(ctx context.Context, dbo dbo.CommercialRate) error {
	q := `insert into commercials (bank, usd_in, usd_out, euro_in, euro_out, rub_in, rub_out) 
	values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := c.db.ExecContext(
		ctx,
		q, dbo.Bank.Name,
		dbo.USD.Buying,
		dbo.USD.Selling,
		dbo.EUR.Buying,
		dbo.EUR.Selling,
		dbo.RUB.Buying,
		dbo.RUB.Selling,
	)
	if err != nil {
		return err
	}

	return err
}
