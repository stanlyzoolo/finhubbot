package schema

import (
	"database/sql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"
)

//go:generate sh -c "$GOPATH/bin/go-bindata -pkg schema -o ./schema_bin.go *.sql"

type Migration struct {
	log *log.Logger
	db  *sql.DB
}

func New(log *log.Logger, db *sql.DB) *Migration {
	// db, err := sql.Open("sqlite3", "./rates.db")
	// if err != nil {
	// 	return nil
	// }

	// if err = db.Ping(); err != nil {
	// 	return nil
	// }

	m := &Migration{
		log: log,
		db:  db,
	}

	if err := m.Up(); err != nil {
		m.log.Error(err)

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
		func(name string) ([]byte, error) {
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
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return err
}
