package main

import (
	"adder/chips"
	"fmt"
)

const (
	Adder       = "Inputs: A=%t, B=%t, Cin=%t | Sum: %t, Carry Out: %t\n"
	MuxTemplate = "Input A: A1=%t, A2=%t, A3=%t, A4=%t\nInput B: B1=%t, B2=%t, B3=%t, B4=%t\nSELECT=%t, STROBE=%t\nOutput: Y1=%t, Y2=%t, Y3=%t, Y4=%t\n"
)

func main() {
	chip := chips.NewFullAdder()

	inputs := map[string]bool{
		"A":   true,
		"B":   true,
		"Cin": false,
	}

	sum, carryOut := chips.RunFullAdder(chip, inputs)

	fmt.Printf(Adder,
		inputs["A"], inputs["B"], inputs["Cin"],
		sum, carryOut,
	)

	inputs = map[string]bool{
		"A":   false,
		"B":   true,
		"Cin": true,
	}
	sum, carryOut = chips.RunFullAdder(chip, inputs)

	fmt.Printf(Adder,
		inputs["A"], inputs["B"], inputs["Cin"],
		sum, carryOut,
	)

	circuit2 := chips.NewMultiplexer()

	// Setup inputs for the multiplexer
	// These would need to be set according to the specific use case
	inputs2 := map[string]bool{
		"A1":     false,
		"A2":     true,
		"A3":     false,
		"A4":     true,
		"B1":     true,
		"B2":     false,
		"B3":     true,
		"B4":     false,
		"SELECT": true, // Assuming SELECT is a single line for simplicity
		"STROBE": true, // Assuming STROBE enables the output when true
	}

	outputs2 := chips.RunMultiplexer(circuit2, inputs2)

	// Printing the inputs and outputs
	fmt.Printf(MuxTemplate,
		inputs2["A1"], inputs2["A2"], inputs2["A3"], inputs2["A4"],
		inputs2["B1"], inputs2["B2"], inputs2["B3"], inputs2["B4"],
		inputs2["SELECT"], inputs2["STROBE"],
		outputs2[0], outputs2[1], outputs2[2], outputs2[3],
	)
}
