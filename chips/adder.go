package chips

import (
	"adder/core"
	"fmt"
)

type FullAdder struct{}

func (fa FullAdder) Create() *core.Circuit {
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
