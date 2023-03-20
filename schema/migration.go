package schema

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/stanlyzoolo/smartLaFamiliaBot/log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
)

var ErrNoChange = fmt.Errorf("no change")

//go:generate sh -c "$GOPATH/bin/go-bindata -pkg schema -o ./schema_bin.go *.sql"
type Migration struct {
	log *log.Logger
	db  *sql.DB
}

func New(log *log.Logger, db *sql.DB) *Migration {
	m := &Migration{
		log: log,
		db:  db,
	}

	err := m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			m.log.Error(err)
		}

		m.log.Info("migration schema has not been changed")

		return nil
	}

	m.log.Info("Migrations up successfully")

	return m
}

func (m *Migration) Up() error {
	dr, err := sqlite3.WithInstance(m.db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	res := bindata.Resource(AssetNames(),
		func(name string) ([]byte, error) { // nolint
			return Asset(name)
		})

	sd, err := bindata.WithInstance(res)
	if err != nil {
		return err
	}

	mn, err := migrate.NewWithInstance("go-bindata", sd, "rates", dr)
	if err != nil {
		return err
	}

	// Если проблемы с накатыванием миграций - смотреть в нейминг *sql в сгенер-ом schema_bin.go
	// https://github.com/golang-migrate/migrate/issues/96?ysclid=lenzdxdzcx20675999
	// Fixed
	err = mn.Up()
	if err != nil {
		return err
	}

	return err
}
