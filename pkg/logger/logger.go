package logger

import (
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

var once sync.Once

var logger *zap.Logger

// Get initializes a zap.Logger instance if it has not been initialized
// already and returns the same instance for subsequent calls.
func Get(level string) (*zap.Logger, error) {
	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	once.Do(func() {
		stdout := zapcore.AddSync(os.Stdout)

		// file := zapcore.AddSync(&lumberjack.Logger{
		// 	Filename:   "logs/app.log",
		// 	MaxSize:    5,
		// 	MaxBackups: 10,
		// 	MaxAge:     14,
		// 	Compress:   true,
		// })

		logLevel := zap.NewAtomicLevelAt(zapLevel)

		productionCfg := zap.NewProductionEncoderConfig()
		productionCfg.TimeKey = "timestamp"
		productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		// developmentCfg := zap.NewDevelopmentEncoderConfig()
		// developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

		consoleEncoder := zapcore.NewJSONEncoder(productionCfg)
		// fileEncoder := zapcore.NewJSONEncoder(productionCfg)

		// log to multiple destinations (console and file)
		// extra fields are added to the JSON output alone
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, stdout, logLevel),
			// zapcore.NewCore(fileEncoder, file, logLevel),
		)

		logger = zap.New(
			core,
			zap.AddCaller(),
		)
	})

	return logger, nil
}

// FromCtx returns the Logger associated with the ctx. If no logger
// is associated, the default logger is returned, unless it is nil
// in which case a disabled logger is returned.
func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	} else if l := logger; l != nil {
		return l
	}

	return zap.NewNop()
}

// WithCtx returns a copy of ctx with the Logger attached.
func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		if lp == l {
			// Do not store same logger.
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, l)
}
