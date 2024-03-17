package chips

import (
	"adder/core"
)

const (
	MuxTemplate = "Input A: A1=%t, A2=%t, A3=%t, A4=%t\nInput B: B1=%t, B2=%t, B3=%t, B4=%t\nSELECT=%t, STROBE=%t\nOutput: Y1=%t, Y2=%t, Y3=%t, Y4=%t\n"
)

func NewMultiplexer() *core.Circuit {
	circuit := core.NewCircuit()

	for i := 1; i <= 4; i++ {
		circuit.AddGate(core.NewGate("NotGate", "NOT"))
		circuit.AddGate(core.NewGate("AndGate", Fmt("AND_A%d", i)))
		circuit.AddGate(core.NewGate("AndGate", Fmt("AND_B%d", i)))
		circuit.AddGate(core.NewGate("OrGate", Fmt("OR_%d", i)))
	}

	for i := 1; i <= 4; i++ {
		circuit.Connect("NOT", "SELECT")
		circuit.Connect(Fmt("AND_A%d", i), Fmt("A%d", i), "SELECT")
		circuit.Connect(Fmt("AND_B%d", i), Fmt("B%d", i), "NOT")
		circuit.Connect(Fmt("Y%d", i), Fmt("OR_%d", i), "STROBE")
		circuit.Connect(Fmt("OR_%d", i), Fmt("AND_A%d", i), Fmt("AND_B%d", i))
	}

	circuit.AddGate(core.NewGate("AndGate", "STROBE_GATE"))
	for i := 1; i <= 4; i++ {
		circuit.Connect("STROBE_GATE", Fmt("Y%d", i), "STROBE")
	}

	return circuit
}

func RunMultiplexer(
	circuit *core.Circuit,
	inputs map[string]bool,
) [4]bool {
	circuit.SetInputs(inputs)

	y1 := circuit.Run("OR_1")
	y2 := circuit.Run("OR_2")
	y3 := circuit.Run("OR_3")
	y4 := circuit.Run("OR_4")

	outputs := [4]bool{y1, y2, y3, y4}

	return outputs
}

func WriteMultiplexer() {
	// // Printing the inputs and outputs
	// fmt.Printf(MuxTemplate,
	// 	inputs2["A1"], inputs2["A2"], inputs2["A3"], inputs2["A4"],
	// 	inputs2["B1"], inputs2["B2"], inputs2["B3"], inputs2["B4"],
	// 	inputs2["SELECT"], inputs2["STROBE"],
	// 	outputs2[0], outputs2[1], outputs2[2], outputs2[3],
	// )
}
