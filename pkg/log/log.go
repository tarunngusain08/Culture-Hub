package log

import (
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger

func New(packageName string) *zap.Logger {
	(&sync.Once{}).Do(func() {
		logger = zap.Must(zap.NewProduction(
			zap.WithCaller(true),
			zap.AddStacktrace(zap.ErrorLevel),
		))
	})

	return logger.With(zap.String("package", packageName))
}

func GetLogger() *zap.Logger {
	return logger
}
