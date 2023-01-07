package logger_test

import (
	"testing"

	"github.com/tae2089/bob-logging/logger"
)

func TestLogging(t *testing.T) {
	logger.Info("123123")
}
