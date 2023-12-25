package main

import (
	"errors"

	"github.com/tae2089/bob-logging/logger"
)

func main() {
	logger.Debug("asdasd")
	logger.Info("123123")
	logger.Error(errors.New("123123"))
}
