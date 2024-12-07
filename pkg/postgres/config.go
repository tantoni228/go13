package postgres

import (
	"fmt"
	"time"
)

type Config struct {
	Host            string        `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
	Port            int           `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
	DB              string        `yaml:"db" env:"POSTGRES_DB" env-default:"postgres"`
	User            string        `yaml:"user" env:"POSTGRES_USER" env-default:"postgres"`
	Password        string        `yaml:"password" env:"POSTGRES_PASSWORD"`
	MaxOpenConns    int           `yaml:"max_open_conns" env:"POSTGRES_MAX_OPEN_CONNS" env-default:"10"`
	MaxIdleConns    int           `yaml:"max_idle_conns" env:"POSTGRES_MAX_IDLE_CONNS" env-default:"5"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" env:"POSTGRES_CONN_MAX_LIFETIME" env-default:"3m"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" env:"POSTGRES_CONN_MAX_IDLE_TIME" env-default:"2m"`
}

func (cfg Config) GetConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
}
