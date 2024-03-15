package main

import (
	"fmt"
)

const (
	Adder = "Inputs: A=%t, B=%t, Cin=%t | Sum: %t, Carry Out: %t\n"
)

func main() {
	circuit := createFullAdderCircuit()

	inputs := map[string]bool{
		"A":   true,
		"B":   true,
		"Cin": false,
	}

	sum, carryOut := RunFullAdder(circuit, inputs)

	fmt.Printf(Adder,
		inputs["A"], inputs["B"], inputs["Cin"],
		sum, carryOut,
	)

	inputs = map[string]bool{
		"A":   false,
		"B":   true,
		"Cin": true,
	}
	sum, carryOut = RunFullAdder(circuit, inputs)

	fmt.Printf(Adder,
		inputs["A"], inputs["B"], inputs["Cin"],
		sum, carryOut,
	)
}
