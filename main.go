package main

import (
	"monolith-accelerator/apackage"
	"monolith-accelerator/utils"

	"go.uber.org/zap"

	"github.com/subosito/gotenv"
)

//Monolith struct
type Monolith struct {
	Logger   *zap.Logger
	APackage *apackage.APackage //example package, A Package
}

//Init initialize Monolith struct
func (m *Monolith) Init() {
	m.Logger = utils.InitLogger()

	m.APackage = &apackage.APackage{}
	m.APackage.Init(m.Logger)
}

//Greeter logger test
func (m *Monolith) Greeter() {
	m.Logger.Info("Hello from Monolith main!")
}

func init() {
	gotenv.Load()
}

func main() {
	mono := &Monolith{}
	mono.Init()
	mono.Greeter()
}
