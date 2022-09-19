package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger  *zap.Logger
	config  zap.Config
	options []zap.Option
}

const (
	KeyTime   = "time" // Match fluentd/fluentbit
	KeyLevel  = "level"
	KeyName   = "name"
	KeyCaller = "caller"
	KeyMsg    = "msg"
	KeyTrace  = "stack"
)

var (
	LevelMap = map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"warn":  zap.WarnLevel,
		"error": zap.ErrorLevel,
		"fatal": zap.FatalLevel,
		"panic": zap.PanicLevel,
	}
)

func New(logger *zap.Logger, config zap.Config, options []zap.Option) *Logger {
	return &Logger{
		logger:  logger,
		config:  config,
		options: options,
	}
}

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
