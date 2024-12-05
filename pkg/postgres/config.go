package postgres

import "fmt"

type Config struct {
	PostgresHost     string `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
	PostgresPort     int    `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
	PostgresDB       string `yaml:"db" env:"POSTGRES_DB" env-default:"postgres"`
	PostgresUser     string `yaml:"user" env:"POSTGRES_USER" env-default:"postgres"`
	PostgresPassword string `yaml:"password" env:"POSTGRES_PASSWORD"`
}

func (cfg Config) GetConnString() string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)
}
