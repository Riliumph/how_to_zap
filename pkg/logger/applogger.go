package logger

import (
	"how_to_zap/pkg/logger/lumberjack"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	app      *Logger
	appLock  sync.Once
	levelMap = map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"warn":  zap.WarnLevel,
		"error": zap.ErrorLevel,
		"fatal": zap.FatalLevel,
		"panic": zap.PanicLevel,
	}
)

func init() {
	appLock.Do(func() {
		// make config
		config := zap.Config{
			Level:    zap.NewAtomicLevelAt(levelMap[os.Getenv("LOG_LEVEL")]),
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				// set default log item
				TimeKey:       keyTime,
				LevelKey:      keyLevel,
				NameKey:       keyName,
				CallerKey:     keyCaller,
				MessageKey:    keyMsg,
				StacktraceKey: keyTrace,
				// set log expression by encoder
				EncodeLevel:    zapcore.CapitalLevelEncoder,   // Log level format: all text is upper case
				EncodeTime:     zapcore.ISO8601TimeEncoder,    // Time format: ISO8601
				EncodeDuration: zapcore.MillisDurationEncoder, // Duration format : seconds for diff calculation
				EncodeCaller:   zapcore.ShortCallerEncoder,    // Stack Trace format: path & line number after the package
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		// make file sink
		fileSink := zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   "/tmp/app.log",
				MaxSize:    100, // megabytes
				MaxBackups: 3,
				MaxAge:     28, //days
			},
		)
		fileEncoder := zapcore.NewJSONEncoder(config.EncoderConfig)
		fileCore := zapcore.NewCore(fileEncoder, fileSink, config.Level)

		// make console sink
		consoleSink := zapcore.AddSync(os.Stdout)
		consoleEncoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
		consoleCore := zapcore.NewCore(consoleEncoder, consoleSink, config.Level)

		// set options
		var options []zap.Option
		options = append(options, zap.AddCallerSkip(1))

		// make core
		core := zapcore.NewTee(
			fileCore,
			consoleCore,
		)

		// make logger
		logger := zap.New(core, options...)
		app = &Logger{
			logger:  logger,
			config:  config,
			options: options,
		}
	})
}

// App Accessor for app
func App() *Logger {
	return app
}
