package utils

import (
	"go.uber.org/zap"
)

//Utils struct
type Utils struct {
}

//InitLogger initialize logger
func InitLogger() *zap.Logger {
	return getConfiguredLogger()
}
