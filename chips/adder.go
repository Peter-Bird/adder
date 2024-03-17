package chips

import (
	"adder/core"
	"fmt"
)

type FullAdder struct{}

func (fa FullAdder) Create() *core.Circuit {
	circuit := core.NewCircuit()

	// Define gates
	gates := []struct {
		Type string
		Name string
	}{
		{"XorGate", "XOR1"},
		{"XorGate", "XOR2"},
		{"AndGate", "AND1"},
		{"AndGate", "AND2"},
		{"OrGate", "OR"},
	}

	// Add gates to the circuit
	for _, gate := range gates {
		circuit.AddGate(core.NewGate(gate.Type, gate.Name))
	}

	// Define connections
	connections := []struct {
		Source string
		Input1 string
		Input2 string
	}{
		{"XOR1", "A", "B"},
		{"XOR2", "XOR1", "Cin"},
		{"AND1", "A", "B"},
		{"AND2", "XOR1", "Cin"},
		{"OR", "AND1", "AND2"},
	}

	// Connect gates within the circuit
	for _, conn := range connections {
		circuit.Connect(conn.Source, conn.Input1, conn.Input2)
	}

	return circuit
}

func (fa FullAdder) Run(
	circuit *core.Circuit,
	inputs map[string]bool,
) interface{} {

	circuit.SetInputs(inputs)

	return map[string]bool{
		"Sum":      circuit.Run("XOR2"),
		"CarryOut": circuit.Run("OR"),
	}
}

func (fa FullAdder) Template() string {
	return "Inputs: A=%t, B=%t, Cin=%t | Sum: %t, Carry Out: %t\n"
}

func (fa FullAdder) Write(
	inputs map[string]bool,
	outputs interface{},
) {

	out, ok := outputs.(map[string]bool)
	if !ok {
		fmt.Println("Error: Invalid outputs format")
		return
	}

	fmt.Printf(
		fa.Template(),
		inputs["A"],
		inputs["B"],
		inputs["Cin"],
		out["Sum"],
		out["CarryOut"],
	)
}
