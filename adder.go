package main

import (
	"fmt"

	"adder/core"
)

func main() {
	// Initialize the circuit
	circuit := core.NewCircuit()

	// Add gates to the circuit
	// Batch add gates to the circuit
	circuit.AddGates([]core.Gate{
		core.NewGate("XorGate", "XOR1"),
		core.NewGate("XorGate", "XOR2"),
		core.NewGate("AndGate", "AND1"),
		core.NewGate("AndGate", "AND2"),
		core.NewGate("OrGate", "OR"),
	})

	circuit.Connect("XOR1", "A", "B")
	circuit.Connect("XOR2", "XOR1", "Cin")
	circuit.Connect("AND1", "A", "B")
	circuit.Connect("AND2", "XOR1", "Cin")
	circuit.Connect("OR", "AND1", "AND2")

	// Define connections (for a full adder circuit)
	circuit.SetInputs(map[string]bool{
		"A":   true,
		"B":   true,
		"Cin": true,
	})

	// Run the circuit to get the sum and carry out
	sum := circuit.Run("XOR2")
	carryOut := circuit.Run("OR")

	// Print the results
	fmt.Printf("Sum: %t, Carry Out: %t\n", sum, carryOut)
}
