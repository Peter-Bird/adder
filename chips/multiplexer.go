package chips

import (
	"adder/core"
	"fmt"
)

type Multiplexer struct{}

func (mux Multiplexer) Create() *core.Circuit {
	circuit := core.NewCircuit()

	// Assuming each "NOT", "AND", and "OR" gate needs a unique identifier
	for i := 1; i <= 4; i++ {
		circuit.AddGates([]core.Gate{
			core.NewGate("NotGate", fmt.Sprintf("NOT%d", i)),
			core.NewGate("AndGate", fmt.Sprintf("AND_A%d", i)),
			core.NewGate("AndGate", fmt.Sprintf("AND_B%d", i)),
			core.NewGate("OrGate", fmt.Sprintf("OR_%d", i)),
		})
	}

	// Correcting connection logic to ensure unique connections per gate
	for i := 1; i <= 4; i++ {
		circuit.Connect(fmt.Sprintf("AND_A%d", i), fmt.Sprintf("A%d", i), "SELECT")
		circuit.Connect(fmt.Sprintf("AND_B%d", i), fmt.Sprintf("B%d", i), fmt.Sprintf("NOT%d", i))
		circuit.Connect(fmt.Sprintf("NOT%d", i), "SELECT")
		circuit.Connect(fmt.Sprintf("OR_%d", i), fmt.Sprintf("AND_A%d", i), fmt.Sprintf("AND_B%d", i))
	}

	// Added logic for STROBE_GATE connections based on the provided logic
	circuit.AddGate(core.NewGate("AndGate", "STROBE_GATE"))
	for i := 1; i <= 4; i++ {
		circuit.Connect("STROBE_GATE", fmt.Sprintf("Y%d", i), "STROBE")
	}

	return circuit
}

func (mux Multiplexer) Run(circuit *core.Circuit, inputs map[string]bool) interface{} {

	circuit.SetInputs(inputs)

	return map[string]bool{
		"Y1": circuit.Run("OR_1"),
		"Y2": circuit.Run("OR_2"),
		"Y3": circuit.Run("OR_3"),
		"Y4": circuit.Run("OR_4"),
	}
}

func (mux Multiplexer) Template() string {
	return `Input A: A1=%t, A2=%t, A3=%t, A4=%t
Input B: B1=%t, B2=%t, B3=%t, B4=%t
SELECT=%t, STROBE=%t
Output: Y1=%t, Y2=%t, Y3=%t, Y4=%t
`
}

func (mux Multiplexer) Write(inputs map[string]bool, outputs interface{}) {
	out, ok := outputs.(map[string]bool)
	if !ok {
		fmt.Println("Error: Invalid outputs format")
		return
	}

	fmt.Printf(
		mux.Template(),
		inputs["A1"], inputs["A2"], inputs["A3"], inputs["A4"],
		inputs["B1"], inputs["B2"], inputs["B3"], inputs["B4"],
		inputs["SELECT"],
		inputs["STROBE"],
		out["Y1"], out["Y2"], out["Y3"], out["Y4"],
	)
}
