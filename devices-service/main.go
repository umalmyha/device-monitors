package main

import (
	"github.com/umalmyha/device-monitors/devices-service/internal/infrastructure/db/mysql"
	"github.com/umalmyha/device-monitors/logger"
	"go.uber.org/zap"
	"log"

	"github.com/umalmyha/device-monitors/devices-service/internal/config"
)

func main() {
	cfg, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}

	_, err = mysql.Connect(cfg.DatabaseConfig.DSN, mysql.Config{
		MaxOpenConns:    cfg.DatabaseConfig.MaxOpenConns,
		MaxIdleConns:    cfg.DatabaseConfig.MaxIdleConns,
		ConnMaxLifetime: cfg.DatabaseConfig.ConnMaxLifetime,
		ConnMaxIdleTime: cfg.DatabaseConfig.ConnMaxIdleTime,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func setupLogger(env string) error {
	var log *zap.Logger
	if env == "production" {
		log = logger.Production()
	}
	log =
}
