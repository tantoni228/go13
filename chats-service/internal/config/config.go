package config

import (
	"go13/chats-service/internal/repo/messages"
	"go13/pkg/postgres"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port        int             `yaml:"server_port" env:"SERVER_PORT" env-default:"8080"`
	LogLevel    string          `yaml:"log_level" env:"LOG_LEVEL" env-default:"INFO"`
	PostgresCfg postgres.Config `yaml:"postgres"`
	MessagesCfg messages.Config `yaml:"messages"`
}

func Get(configPath string) (Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
