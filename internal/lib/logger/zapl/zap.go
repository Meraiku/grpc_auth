package zapl

import (
	"go.uber.org/zap"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(env string) *zap.Logger {
	var log *zap.Logger

	switch env {
	case envLocal:
		log, _ = zap.NewDevelopment()
	case envDev:
		log, _ = zap.NewDevelopment()
	case envProd:
		log, _ = zap.NewProduction()
	}

	return log
}
