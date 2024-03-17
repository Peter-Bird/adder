package chips

import (
	"adder/core"
	"fmt"
)

type Multiplexer struct{}

func (mux Multiplexer) Create() *core.Circuit {
	circuit := core.NewCircuit()

	gates := []struct {
		Type string
		Name string
	}{
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
	}

	for _, gate := range gates {
		circuit.AddGate(core.NewGate(gate.Type, gate.Name))
	}

	connections := [][]string{
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
	}

	for _, conn := range connections {
		source := conn[0]
		inputs := conn[1:]
		circuit.Connect(source, inputs...)
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
