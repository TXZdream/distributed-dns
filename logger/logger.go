package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	logger, _ := zap.NewDevelopment()
	Logger = logger
}
