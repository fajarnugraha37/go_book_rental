package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func initialize() {
	if logger == nil {
		// encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg := zap.NewDevelopmentEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		config := zap.Config{
			Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          "console", // json
			EncoderConfig:     encoderCfg,
			OutputPaths: []string{
				"stderr",
			},
			ErrorOutputPaths: []string{
				"stderr",
			},
			// InitialFields: map[string]interface{}{
			// 	"pid": os.Getpid(),
			// },
		}
		logger = zap.Must(config.Build())
		sugar = logger.Sugar()
	}
	if sugar == nil {
		sugar = logger.Sugar()
	}
}

func GetLogger() *zap.Logger {
	initialize()

	return logger
}

func GetSugaredLogger() *zap.SugaredLogger {
	initialize()

	return sugar
}
