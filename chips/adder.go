package chips

import (
	"adder/core"
)

type FullAdder struct {
	GateType    string
	Gates       Gates
	Connections Connections
	InPorts     Ports
	OutPorts    Ports
}

func NewFA() *FullAdder {
	return &FullAdder{
		GateType: "FullAdder",
		Gates: Gates{
			{"XorGate", "XOR1"},
			{"XorGate", "XOR2"},
			{"AndGate", "AND1"},
			{"AndGate", "AND2"},
			{"OrGate", "OR"},
			{"EquGate", "Sum"},
			{"EquGate", "Cout"},
		},
		Connections: Connections{
			{"XOR1", "A", "B"},
			{"XOR2", "XOR1", "Cin"},
			{"AND1", "A", "B"},
			{"AND2", "XOR1", "Cin"},
			{"OR", "AND1", "AND2"},
			{"Sum", "XOR2"},
			{"Cout", "OR"},
		},
		InPorts: Ports{
			"A",
			"B",
			"Cin",
		},
		OutPorts: Ports{
			"Sum",
			"Cout",
		},
	}
}

func (fa FullAdder) Create() *core.Circuit {
	circuit := core.NewCircuit()

	// Add gates to circuit
	for _, gate := range fa.Gates {
		circuit.AddGate(core.NewGate(gate.Type, gate.Name))
	}

	// Connect gates within the circuit
	for _, conn := range fa.Connections {
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

	outputs := make(map[string]bool)

	for _, output := range fa.OutPorts {
		outputs[output] = circuit.Run(output)
	}

	return outputs
}

func (fa FullAdder) Write(inputs, outputs map[string]bool) {
	Write(inputs, outputs)
}
