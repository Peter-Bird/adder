package chips

import "adder/core"

type FullAdder struct {
	core.Chip
}

func NewFA() *FullAdder {
	return &FullAdder{
		Chip: core.Chip{
			GateType: "FullAdder",
			Gates: core.Gates{
				{"XorGate", "XOR1"},
				{"XorGate", "XOR2"},
				{"AndGate", "AND1"},
				{"AndGate", "AND2"},
				{"OrGate", "OR"},
				{"EquGate", "Sum"},
				{"EquGate", "Cout"},
			},
			Connections: core.Connections{
				{"XOR1", "A", "B"},
				{"XOR2", "XOR1", "Cin"},
				{"AND1", "A", "B"},
				{"AND2", "XOR1", "Cin"},
				{"OR", "AND1", "AND2"},
				{"Sum", "XOR2"},
				{"Cout", "OR"},
			},
			InPorts: core.PortList{
				"A",
				"B",
				"Cin",
			},
			OutPorts: core.PortList{
				"Sum",
				"Cout",
			},
		},
	}
}
