package log_zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Configuração do logger
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig = encoderConfig

	// Criação do logger
	logger, err := loggerConfig.Build()
	if err != nil {
		panic("Erro ao criar logger: " + err.Error())
	}
	return logger
}
