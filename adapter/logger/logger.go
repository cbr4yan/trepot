package logger

import (
	"github.com/cbr4yan/trepot/config"
	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}

func Setup(cfg *config.Config) error {
	if cfg.IsProduction() {
		logger, err := zap.NewProduction()
		if err != nil {
			return err
		}
		zap.ReplaceGlobals(logger)
	}
	return nil
}
