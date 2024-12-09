package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var (
    Log = loggerBuild()
)

// Имплементация интерфейса логгера
type Logger struct {
	logger	*zap.Logger
}

// loggerBuild инициализирует и возвращает экземпляр логгера Zap
func loggerBuild() *Logger {
    config := zap.NewDevelopmentConfig()
    config.OutputPaths = []string{"../../log/log.log", "stdout"}

    logger, err := config.Build()

    if err != nil {
        panic(fmt.Sprintf("failed to configure logger: %v", err))
    }

    return &Logger{
        logger: logger,
    }
}

func (l Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}