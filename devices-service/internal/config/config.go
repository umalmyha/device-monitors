package config

import (
	"github.com/caarlos0/env/v8"
	"time"
)

type DatabaseConfig struct {
	DSN             string        `env:"MYSQL_URL,notEmpty"`
	MaxOpenConns    int           `env:"POSTGRES_MAX_OPEN_CONNS"`
	MaxIdleConns    int           `env:"POSTGRES_MAX_IDLE_CONNS" envDefault:"2"`
	ConnMaxLifetime time.Duration `env:"POSTGRES_CONN_MAX_LIFETIME"`
	ConnMaxIdleTime time.Duration `env:"POSTGRES_CONN_IDLE_TIME"`
}

type Config struct {
	Environment    string `env:"ENVIRONMENT" envDefault:"development"`
	Port           int    `env:"PORT" envDefault:"8080"`
	DatabaseConfig DatabaseConfig
}

func Build() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
