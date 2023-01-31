package log

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func New() (*Logger, error) {
	instance, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("can't initialize logger: %w", err)
	}

	return &Logger{
		SugaredLogger: instance.Sugar(),
	}, nil
}
