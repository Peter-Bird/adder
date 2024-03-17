package main

import (
	"adder/chips"
)

func main() {
	fa()
	mux()
}

func fa() {
	fa := chips.FullAdder{}
	chip := fa.Create()

	inputs := map[string]bool{
		"A": true, "B": true, "Cin": true,
	}

	outputs := fa.Run(chip, inputs).(map[string]bool)

	fa.Write(inputs, outputs)
}

func mux() {
	mux := chips.Multiplexer{}
	chip := mux.Create()

	inputs := map[string]bool{
		"A1": false, "A2": true, "A3": false, "A4": true,
		"B1": true, "B2": false, "B3": true, "B4": false,
		"SELECT": true,
		"STROBE": true,
	}

	outputs := mux.Run(chip, inputs).(map[string]bool)

	mux.Write(inputs, outputs)
}
