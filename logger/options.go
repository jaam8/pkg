package logger

import (
	"time"

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
			switch f.Type {
			case zapcore.StringType:
				c.InitialFields[f.Key] = f.String
			case zapcore.Int64Type, zapcore.Int32Type, zapcore.Int16Type, zapcore.Int8Type, zapcore.Uint64Type, zapcore.Uint32Type, zapcore.Uint16Type, zapcore.Uint8Type:
				c.InitialFields[f.Key] = f.Integer
			case zapcore.BoolType:
				c.InitialFields[f.Key] = f.Integer == 1
			case zapcore.Float64Type, zapcore.Float32Type:
				c.InitialFields[f.Key] = f.Integer
			case zapcore.DurationType:
				c.InitialFields[f.Key] = time.Duration(f.Integer)
			case zapcore.TimeType:
				c.InitialFields[f.Key] = f.Interface
			default:
				c.InitialFields[f.Key] = f.Interface
			}
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
