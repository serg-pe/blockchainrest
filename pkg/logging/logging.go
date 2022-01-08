package logging

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	return logger
}

func init() {
	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logConfig.EncoderConfig.CallerKey = "caller"
	configuredLogger, err := logConfig.Build()
	if err != nil {
		log.Fatalln("Can't initialize logger")
	}
	logger = configuredLogger
}
