package logger

import "go.uber.org/zap/zapcore"

func ParseLevel(s string) (zapcore.Level, error) {
	var lvl zapcore.Level
	if err := lvl.Set(s); err != nil {
		return zapcore.InfoLevel, err
	}
	return lvl, nil
}
