package chips

import (
	"adder/core"
)

type Multiplexer struct {
	GateType    string
	Gates       Gates
	Connections Connections
	InPorts     Ports
	OutPorts    Ports
}

func NewMux() *Multiplexer {
	return &Multiplexer{
		GateType: "Multiplexer",
		Gates: Gates{
			{"NotGate", "NOT1"},
			{"NotGate", "NOT2"},
			{"NotGate", "NOT3"},
			{"NotGate", "NOT4"},
			{"AndGate", "AND_A1"},
			{"AndGate", "AND_A2"},
			{"AndGate", "AND_A3"},
			{"AndGate", "AND_A4"},
			{"AndGate", "AND_B1"},
			{"AndGate", "AND_B2"},
			{"AndGate", "AND_B3"},
			{"AndGate", "AND_B4"},
			{"OrGate", "OR_1"},
			{"OrGate", "OR_2"},
			{"OrGate", "OR_3"},
			{"OrGate", "OR_4"},
			{"AndGate", "STROBE_GATE"},
			{"EquGate", "Y1"},
			{"EquGate", "Y2"},
			{"EquGate", "Y3"},
			{"EquGate", "Y4"},
		},
		Connections: Connections{
			{"AND_A1", "A1", "SELECT"},
			{"AND_B1", "B1", "NOT1"},
			{"NOT1", "SELECT"},
			{"OR_1", "AND_A1", "AND_B1"},
			{"AND_A2", "A2", "SELECT"},
			{"AND_B2", "B2", "NOT2"},
			{"NOT2", "SELECT"},
			{"OR_2", "AND_A2", "AND_B2"},
			{"AND_A3", "A3", "SELECT"},
			{"AND_B3", "B3", "NOT3"},
			{"NOT3", "SELECT"},
			{"OR_3", "AND_A3", "AND_B3"},
			{"AND_A4", "A4", "SELECT"},
			{"AND_B4", "B4", "NOT4"},
			{"NOT4", "SELECT"},
			{"OR_4", "AND_A4", "AND_B4"},
			{"STROBE_GATE", "OR_1", "Y1"},
			{"STROBE_GATE", "OR_2", "Y2"},
			{"STROBE_GATE", "OR_3", "Y3"},
			{"STROBE_GATE", "OR_4", "Y4"},
			{"STROBE_GATE", "STROBE"},
			{"Y1", "OR_1"},
			{"Y2", "OR_2"},
			{"Y3", "OR_3"},
			{"Y4", "OR_4"},
		},
		InPorts: Ports{
			"A1",
			"A2",
			"A3",
			"A4",
			"B1",
			"B2",
			"B3",
			"B4",
			"SELECT",
			"STROBE",
		},
		OutPorts: Ports{
			"Y1",
			"Y2",
			"Y3",
			"Y4",
		},
	}
}

func (mux Multiplexer) Create() *core.Circuit {
	circuit := core.NewCircuit()

	for _, gate := range mux.Gates {
		circuit.AddGate(core.NewGate(gate.Type, gate.Name))
	}

	// Connect gates within the circuit
	for _, conn := range mux.Connections {
		source := conn[0]
		inputs := conn[1:]
		circuit.Connect(source, inputs...)
	}

	return circuit
}

func (mux Multiplexer) Run(circuit *core.Circuit, inputs map[string]bool) interface{} {

	circuit.SetInputs(inputs)

	outputs := make(map[string]bool)

	for _, output := range mux.OutPorts {
		outputs[output] = circuit.Run(output)
	}

	return outputs
}

func (mux Multiplexer) Write(
	inputs map[string]bool,
	outputs map[string]bool,
) {

	Write(inputs, outputs)
}
