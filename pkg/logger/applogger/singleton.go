package applogger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"how_to_zap/pkg/logger"
	"how_to_zap/pkg/logger/lumberjack"
)

var (
	app     *logger.Logger
	appLock sync.Once
)

func init() {
	appLock.Do(func() {
		// make config
		config := zap.Config{
			Level:    zap.NewAtomicLevelAt(logger.LevelMap[os.Getenv("LOG_LEVEL")]),
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				// set default log item
				TimeKey:       logger.KeyTime,
				LevelKey:      logger.KeyLevel,
				NameKey:       logger.KeyName,
				CallerKey:     logger.KeyCaller,
				MessageKey:    logger.KeyMsg,
				StacktraceKey: logger.KeyTrace,
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
		instance := zap.New(core, options...)
		app = logger.New(instance, config, options)
	})
}

// App Accessor for app
func App() *logger.Logger {
	return app
}
