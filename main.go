package main

import (
	"adder/chips"
	"adder/core"
)

func main() {
	fa()
	mux()
}

func fa() {
	fa := chips.NewFA()
	chip := fa.Create()

	inputs := core.Ports{
		"A": true, "B": true, "Cin": false,
	}

	outputs := fa.Run(chip, inputs)

	fa.Write(inputs, outputs)
}

func mux() {
	mux := chips.NewMux()
	chip := mux.Create()

	inputs := core.Ports{
		"A1": false, "A2": true, "A3": false, "A4": true,
		"B1": true, "B2": false, "B3": true, "B4": false,
		"SELECT": true,
		"STROBE": true,
	}

	outputs := mux.Run(chip, inputs)

	mux.Write(inputs, outputs)
}
