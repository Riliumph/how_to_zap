package main

import (
	l "how_to_zap/pkg/logger/applogger"

	"go.uber.org/zap"
)

func main() {
	l.App().With(zap.Any("request_id", "foobar"))
	l.App().Debug("show parameter", zap.Any("param1", 1))
}
