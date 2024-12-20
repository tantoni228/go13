package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

func Get(ctx context.Context, cfg Config) (*Postgres, error) {
	db, err := sqlx.Open("postgres", cfg.GetConnString())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}
