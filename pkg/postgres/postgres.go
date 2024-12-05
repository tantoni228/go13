package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

func Get(cfg Config) (*Postgres, error) {
	db, err := sqlx.Connect("postgres", cfg.GetConnString())
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}
