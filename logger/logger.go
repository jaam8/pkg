package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type requestIDKey struct{}
type loggerKey struct{}

type Logger struct {
	l *zap.Logger
}

// New creates a new logger and attaches it to the provided context.
// It accepts optional configuration options.
// Example usage:
//
//	ctx, err := logger.New(context.Background(), logger.WithLevel(lvl))
//	if err != nil {
//	// handle error
//	}
func New(ctx context.Context, opts ...Option) (context.Context, error) {
	cfg := zap.NewProductionConfig()

	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	cfg.Level.SetLevel(zap.InfoLevel)

	for _, opt := range opts {
		opt(&cfg)
	}

	l, err := cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, loggerKey{}, Logger{l: l})

	return ctx, nil
}

func FromCtx(ctx context.Context) *Logger {
	if l, ok := ctx.Value(loggerKey{}).(*Logger); ok {
		return l
	}

	return nil
}

func WithRID(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, rid)
}

func tryAddRID(ctx context.Context, fields []zap.Field) []zap.Field {
	if ctx.Value(requestIDKey{}) != nil {
		if rid, ok := ctx.Value(requestIDKey{}).(string); ok {
			fields = append(fields, zap.String("rid", rid))
		}
	}
	return fields
}
