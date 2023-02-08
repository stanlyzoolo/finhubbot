package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/stanlyzoolo/smartLaFamiliaBot/banks/commercial"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"

	_ "github.com/mattn/go-sqlite3"
)

type Commercials interface {
	Create(ctx context.Context, r commercial.Rate) error
}

type commercials struct {
	db  *sql.DB
	log *log.Logger
}

func NewCommercials(log *log.Logger) Commercials {
	var err error

	db, err := sql.Open("sqlite3", "./rates.db")
	if err != nil {
		return nil
	}

	if err = db.Ping(); err != nil {
		return nil
	}

	return &commercials{
		db:  db,
		log: log,
	}
}

func (c *commercials) Create(ctx context.Context, r commercial.Rate) error {
	q := `insert into commercials (bank, usd_in, usd_out, euro_in, euro_out, rub_in, rub_out, conv_usd_to_euro_in, conv_usd_to_euro_out, date) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := c.db.ExecContext(
		ctx,
		q, r.Bank.Name,
		r.USD.Buying,
		r.USD.Selling,
		r.EUR.Buying,
		r.EUR.Selling,
		r.RUB.Buying,
		r.RUB.Selling,
		r.ConvFromEURtoUSD.Buying,
		r.ConvFromEURtoUSD.Selling,
		time.Now(),
	)
	if err != nil {
		return err
	}

	return err
}
