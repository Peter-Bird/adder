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
	connections := [][]string{
		{"XOR1", "A", "B"},
		{"XOR2", "XOR1", "Cin"},
		{"AND1", "A", "B"},
		{"AND2", "XOR1", "Cin"},
		{"OR", "AND1", "AND2"},
	}

	// Connect gates within the circuit
	for _, conn := range connections {
		source := conn[0]
		inputs := conn[1:]
		circuit.Connect(source, inputs...)
	}

	return circuit
}

func (fa FullAdder) Run(
	circuit *core.Circuit,
	inputs map[string]bool,
) interface{} {

	circuit.SetInputs(inputs)

	outMap := map[string]string{
		"Sum":      "XOR2",
		"CarryOut": "OR",
	}

	outputs := make(map[string]bool)

	for outputName, gateName := range outMap {
		outputs[outputName] = circuit.Run(gateName)
	}

	return outputs
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
