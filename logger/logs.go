package logger

import (
	"context"

	"go.uber.org/zap"
)

// With adds the given fields to the logger in the context and returns a new context.
//
// Example usage:
//
//	ctx = logger.With(ctx, zap.String("key", "value"))
func With(ctx context.Context, fields ...zap.Field) context.Context {
	l := FromCtx(ctx)
	if l == nil || l.l == nil {
		return ctx
	}

	nl := l.l.With(fields...)
	return context.WithValue(ctx, loggerKey{}, &Logger{l: nl})
}

// Info logs an info level message with the given fields.
//
// Example usage:
//
//	logger.Info(ctx, "this is an info message", zap.String("key", "value"))
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = tryAddRID(ctx, fields)
	if l := FromCtx(ctx); l != nil {
		l.l.Info(msg, fields...)
	}
}

// Debug logs a debug level message with the given fields.
//
// Example usage:
//
//	logger.Debug(ctx, "this is a debug message", zap.String("key", "value"))
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = tryAddRID(ctx, fields)
	if l := FromCtx(ctx); l != nil {
		l.l.Debug(msg, fields...)
	}
}

// Warn logs a warning level message with the given fields.
//
// Example usage:
//
//	logger.Warn(ctx, "this is a warning message", zap.String("key", "value"))
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = tryAddRID(ctx, fields)
	if l := FromCtx(ctx); l != nil {
		l.l.Warn(msg, fields...)
	}
}

// Error logs an error level message with the given fields.
//
// Example usage:
//
//	logger.Error(ctx, "this is an error message", zap.String("key", "value"))
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = tryAddRID(ctx, fields)
	if l := FromCtx(ctx); l != nil {
		l.l.Error(msg, fields...)
	}
}
