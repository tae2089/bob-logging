package main

import (
	"errors"

	"github.com/tae2089/bob-logging/logger"
	"go.uber.org/zap"
)

func main() {
	logger.Debug("asdasd")
	logger.Info("123123")
	logger.Error("error", zap.Error(errors.New("123123")))
}
