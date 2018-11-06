package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func New() *zap.Logger {
	logger, err := zap.NewProduction(zap.AddCaller())
	if err != nil {
		panic(errors.Wrap(err, "logger creation"))
	}
	return logger
}
