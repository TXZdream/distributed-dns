package logger

import "go.uber.org/zap"

import "log"

var Logger *zap.Logger

func init() {
	// logger, _ := zap.NewDevelopment()
	cfg := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: zap.NewDevelopmentConfig().Encoding,
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalln(err)
	}
	Logger = logger
}
