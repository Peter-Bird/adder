package chips

type FullAdder struct {
	Chip
}

func NewFA() *FullAdder {
	return &FullAdder{
		Chip: Chip{
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
			InPorts: PortList{
				"A",
				"B",
				"Cin",
			},
			OutPorts: PortList{
				"Sum",
				"Cout",
			},
		},
	}
}
