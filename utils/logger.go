package utils

import (
	"log"
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLogEncodeLevel() (lvlEnc zapcore.LevelEncoder) {
	switch os.Getenv("LOG_ENCODE_LEVEL") {
	case "LowercaseLevelEncoder":
		lvlEnc = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		lvlEnc = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder":
		lvlEnc = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		lvlEnc = zapcore.CapitalColorLevelEncoder
	default:
		lvlEnc = zapcore.CapitalLevelEncoder
	}

	return
}

func getConfiguredLogger() *zap.Logger {
	cfg := &zap.Config{
		//OutputPaths:      []string{"stdout", "./logs/logs"},
		//OutputPaths:      []string{"./logs/logs"},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey: "level",

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			//CallerKey:    "caller",
			//EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	loglvl, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	cfg.Level = zap.NewAtomicLevelAt(zapcore.Level(loglvl))
	cfg.Encoding = os.Getenv("LOG_ENDCODING")
	cfg.EncoderConfig.EncodeLevel = getLogEncodeLevel()

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")
	return logger
}
