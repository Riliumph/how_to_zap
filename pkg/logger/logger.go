package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger  *zap.Logger
	config  zap.Config
	options []zap.Option
}

const (
	keyTime   = "time" // Match fluentd/fluentbit
	keyLevel  = "level"
	keyName   = "name"
	keyCaller = "caller"
	keyMsg    = "msg"
	keyTrace  = "stack"
)

// With Add item
func (l *Logger) With(fields ...zap.Field) {
	if l == nil || l.logger == nil {
		// avoid unexpected nil access
		// ex) test code
		return
	}
	l.logger = l.logger.With(fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	if len(fields) != 0 {
		l.With(zap.Namespace("details"))
	}
	defer l.logger.Sync()
	l.logger.Debug(msg, fields...)
}

func (l *Logger) info(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	if len(fields) != 0 {
		l.With(zap.Namespace("details"))
	}
	defer l.logger.Sync()
	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	if len(fields) != 0 {
		l.With(zap.Namespace("details"))
	}
	defer l.logger.Sync()
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	if len(fields) != 0 {
		l.With(zap.Namespace("details"))
	}
	defer l.logger.Sync()
	l.logger.Error(msg, fields...)
}
