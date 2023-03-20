package repo

import (
	"context"
	"database/sql"

	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/dbo"

	_ "github.com/mattn/go-sqlite3" // nolint
)

type NatBankRB interface {
	Create(ctx context.Context, dbo dbo.NatBankRate) error
}

type natBankRB struct {
	db  *sql.DB
	log *log.Logger
}

func NewNBRB(db *sql.DB, log *log.Logger) NatBankRB {
	return &natBankRB{
		db:  db,
		log: log,
	}
}

func (n *natBankRB) Create(ctx context.Context, dbo dbo.NatBankRate) error {
	q := `insert into nat_bank (nat_id, abbreviation, name, scale, official_rate) 
	values ($1, $2, $3, $4, $5)`

	_, err := n.db.ExecContext(ctx, q, dbo.ID, dbo.Abbreviation.String, dbo.Name.String, dbo.Scale, dbo.OfficialRate)
	if err != nil {
		return err
	}

	return err
}
