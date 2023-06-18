package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)
	// consoleDebugging := zapcore.Lock(os.Stdout)

	consoleErrors := zapcore.Lock(os.Stderr)
	writeStdout := zapcore.AddSync(CustomWriter{})
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.NewMultiWriteSyncer(writeStdout), zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, consoleErrors, zapcore.ErrorLevel),
	)

	log = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)

}

func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

func Error(err error, fields ...zapcore.Field) {
	log.Error(err.Error(), fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	log.Warn(msg, fields...)
}
