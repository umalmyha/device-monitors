package main

import (
	"fmt"
	"github.com/umalmyha/device-monitors/devices-service/internal/query"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/umalmyha/device-monitors/devices-service/internal/config"
	"github.com/umalmyha/device-monitors/devices-service/internal/handler"
	"github.com/umalmyha/device-monitors/devices-service/internal/infrastructure/db/mysql"
	"github.com/umalmyha/device-monitors/devices-service/internal/infrastructure/logger"
	"github.com/umalmyha/device-monitors/devices-service/internal/middleware"
	"github.com/umalmyha/device-monitors/devices-service/internal/model"
	"github.com/umalmyha/device-monitors/devices-service/internal/repository"
	"github.com/umalmyha/device-monitors/devices-service/internal/service"
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

	query.SetDefault(db)

	if err = db.AutoMigrate(&model.Device{}); err != nil {
		zap.S().Fatal(err)
	}

	// repo
	deviceRepo := repository.NewMySqlDeviceRepository(db)

	// service
	deviceSrv := service.NewDeviceService(deviceRepo)

	// handler
	deviceHandler := handler.NewDeviceHandler(deviceSrv)

	r := gin.New()
	r.Use(middleware.ErrorMiddleware, gin.Recovery())
	deviceGrp := r.Group("/devices")
	deviceGrp.GET("/", deviceHandler.FindAll)
	deviceGrp.GET("/:id", deviceHandler.FindByID)
	deviceGrp.POST("/", deviceHandler.Create)
	deviceGrp.PUT("/:id", deviceHandler.Update)
	deviceGrp.DELETE("/:id", deviceHandler.Delete)

	if err = r.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		zap.S().Fatal(err)
	}
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
