package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func Log() *zap.SugaredLogger {
	if log == nil {
		log = zap.S()
	}
	return log
}

func SetGlobal(logger *zap.Logger) {
	zap.ReplaceGlobals(logger)
	log = zap.S()
}

func Production(opts ...zap.Option) (*zap.Logger, error) {
	return config().Build(opts...)
}

func Development(opts ...zap.Option) (*zap.Logger, error) {
	cfg := config()
	cfg.DisableStacktrace = true
	return cfg.Build(opts...)
}

func config() zap.Config {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{"stdout"}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return cfg
}
