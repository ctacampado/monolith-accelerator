package main

import (
	"testing"
)

func TestGreeter(t *testing.T) {
	mono := &Monolith{}
	mono.Init()

	mono.Greeter()
	mono.APackage.Greeter()
}
