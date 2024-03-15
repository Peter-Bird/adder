package main

import (
	"adder/core"
)

func createFullAdderCircuit() *core.Circuit {
	circuit := core.NewCircuit()

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

	return circuit
}

func RunFullAdder(
	circuit *core.Circuit,
	inputs map[string]bool,
) (sum bool, carryOut bool) {

	circuit.SetInputs(inputs)

	sum = circuit.Run("XOR2")
	carryOut = circuit.Run("OR")

	return
}
