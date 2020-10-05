package adapters

import (
	"github.com/golang-clean-architecture/core/depedencies"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerAdapter struct {
	logger    *zap.SugaredLogger
	ZapLogger *zap.Logger
}

func NewLoggerAdapter() LoggerAdapter {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.LevelKey = "log_level"
	encoderConfig.MessageKey = "message"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.TimeKey = "timestamp_app"

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig:    encoderConfig,
		DisableCaller:    true,
	}

	ZapLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	logger := ZapLogger.Sugar()
	defer logger.Sync()
	return LoggerAdapter{
		logger:    logger,
		ZapLogger: ZapLogger,
	}
}

func addEvent(event ...depedencies.Event) zap.Field {
	if len(event) > 0 {
		return zap.Any("event", event[0])
	}
	return zap.Any("event", nil)
}

func (adapter LoggerAdapter) Debug(msg string, event ...depedencies.Event) {
	adapter.logger.Debugw(msg, addEvent(event...))
}

func (adapter LoggerAdapter) Info(msg string, event ...depedencies.Event) {
	adapter.logger.Infow(msg, addEvent(event...))
}

func (adapter LoggerAdapter) Warn(msg string, event ...depedencies.Event) {
	adapter.logger.Warnw(msg, addEvent(event...))
}

func (adapter LoggerAdapter) Error(msg string, err error, event ...depedencies.Event) {
	adapter.logger.Errorw(msg, zap.Error(err), addEvent(event...))
}

func (adapter LoggerAdapter) Fatal(msg string, err error, event ...depedencies.Event) {
	adapter.logger.Fatalw(msg, zap.Error(err), addEvent(event...))
}
