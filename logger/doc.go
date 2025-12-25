// Package logger
// based on uber-go/zap logger
//
// Example usage:
//
//	lvl, err := logger.ParseLevel(cfg.LogLevel)
//	if err != nil {
//	    // handle error
//	}
//	ctx, err := logger.New(context.Background(),
//		logger.WithLevel(lvl),
//		logger.WithFields(
//			zap.String("service", "my-service"),
//			zap.String("env", cfg.Env),
//		)
//	if err != nil {
//	    // handle error
//	}
//
//	logger.Info(ctx, "logger initialized")
package logger
