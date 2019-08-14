package apackage

import (
	"go.uber.org/zap"
)

type APackage struct {
	Logger *zap.Logger
}

func (ap *APackage) Init(logger *zap.Logger) {
	ap.Logger = logger
}

func (ap *APackage) Greeter() {
	ap.Logger.Info("Hello from Monolith apackage!")
}
