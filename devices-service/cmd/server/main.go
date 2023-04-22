package main

import (
	"github.com/gin-gonic/gin"
	"github.com/umalmyha/device-monitors/devices-service/internal/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"

	"github.com/umalmyha/device-monitors/devices-service/internal/config"
	"github.com/umalmyha/device-monitors/devices-service/internal/infrastructure/db/mysql"
	"github.com/umalmyha/device-monitors/devices-service/internal/infrastructure/logger"
)

func main() {
	cfg, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}

	zapLogger, err := zapLogger(cfg.Environment)
	if err != nil {
		log.Fatal(err)
	}
	defer zapLogger.Sync()

	zap.ReplaceGlobals(zapLogger)

	db, err := mysql.Connect(cfg.DatabaseConfig.DSN, mysql.Config{
		MaxOpenConns:    cfg.DatabaseConfig.MaxOpenConns,
		MaxIdleConns:    cfg.DatabaseConfig.MaxIdleConns,
		ConnMaxLifetime: cfg.DatabaseConfig.ConnMaxLifetime,
		ConnMaxIdleTime: cfg.DatabaseConfig.ConnMaxIdleTime,
	})
	if err != nil {
		zap.S().Fatal(err)
	}

	if err = db.AutoMigrate(&model.Device{}); err != nil {
		zap.S().Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	dvcRoute := r.Group("/devices")
	dvcRoute.GET("/devices")
}

func zapLogger(env string) (*zap.Logger, error) {
	opts := zap.Fields(zap.Field{
		Key:    "service",
		Type:   zapcore.StringType,
		String: "devices service",
	})

	log, err := logger.ForEnv(env, opts)
	if err != nil {
		return nil, err
	}
	return log, nil
}
