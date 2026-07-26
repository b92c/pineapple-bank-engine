package logger

import "go.uber.org/zap"

func NewLogger(service string) *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		return zap.NewNop()
	}
	if service != "" {
		return l.Named(service)
	}
	return l
}
