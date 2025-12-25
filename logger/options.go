package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(*zap.Config)

// WithLevel sets the logging level.
func WithLevel(level zapcore.Level) Option {
	return func(c *zap.Config) {
		c.Level.SetLevel(level)
	}
}

// WithTimeEncoder sets the time encoder for the logger.
func WithTimeEncoder(enc zapcore.TimeEncoder) Option {
	return func(c *zap.Config) {
		c.EncoderConfig.EncodeTime = enc
	}
}

// WithCaller enables/disables caller information in the logs.
func WithCaller(enabled bool) Option {
	return func(c *zap.Config) {
		c.DisableCaller = !enabled
	}
}

// WithStacktrace enables/disables stacktrace in the logs.
func WithStacktrace(enabled bool) Option {
	return func(c *zap.Config) {
		c.DisableStacktrace = !enabled
	}
}

// WithFields adds initial fields to the logger.
func WithFields(fields ...zap.Field) Option {
	return func(c *zap.Config) {
		if c.InitialFields == nil {
			c.InitialFields = make(map[string]interface{})
		}
		for _, f := range fields {
			c.InitialFields[f.Key] = f.Interface
		}
	}
}

// WithOutput sets the output paths for the logger.
func WithOutput(paths ...string) Option {
	return func(c *zap.Config) {
		c.OutputPaths = paths
		c.ErrorOutputPaths = paths
	}
}

// WithTimeKey sets the key used for the time field in the logs.
func WithTimeKey(key string) Option {
	return func(c *zap.Config) {
		c.EncoderConfig.TimeKey = key
	}
}
