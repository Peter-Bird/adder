package core

type Gate interface {
	GetLabel() string
	Exec(*Circuit) bool
}

var makeGate = map[string]func(string) Gate{
	"XorGate":  func(l string) Gate { return &XorGate{BasicGate: BasicGate{Label: l}} },
	"AndGate":  func(l string) Gate { return &AndGate{BasicGate: BasicGate{Label: l}} },
	"OrGate":   func(l string) Gate { return &OrGate{BasicGate: BasicGate{Label: l}} },
	"NotGate":  func(l string) Gate { return &NotGate{BasicGate: BasicGate{Label: l}} },
	"NandGate": func(l string) Gate { return &NandGate{BasicGate: BasicGate{Label: l}} },
	"NorGate":  func(l string) Gate { return &NorGate{BasicGate: BasicGate{Label: l}} },
	"XnorGate": func(l string) Gate { return &XnorGate{BasicGate: BasicGate{Label: l}} },
	"EquGate":  func(l string) Gate { return &XnorGate{BasicGate: BasicGate{Label: l}} },
}

func NewGate(gateType, label string) Gate {
	if factory, exists := makeGate[gateType]; exists {
		return factory(label)
	}
	return nil
}

type BasicGate struct {
	Label string
}

func (g *BasicGate) GetLabel() string {
	return g.Label
}

type XorGate struct{ BasicGate }
type AndGate struct{ BasicGate }
type OrGate struct{ BasicGate }
type NotGate struct{ BasicGate }
type NandGate struct{ BasicGate }
type NorGate struct{ BasicGate }
type XnorGate struct{ BasicGate }
type EquGate struct{ BasicGate }

// XorGate.Exec: XOR is true if and only
// if there's an odd number of true inputs.
func (g *XorGate) Exec(circuit *Circuit) bool {
	inputs := circuit.GetInputs(g.Label)
	trueCount := 0

	for _, input := range inputs {
		if input {
			trueCount++
		}
	}

	return trueCount%2 == 1
}

// AndGate.Exec: Returns false if any input is false
func (g *AndGate) Exec(circuit *Circuit) bool {
	inputs := circuit.GetInputs(g.Label)

	for _, input := range inputs {
		if !input {
			return false
		}
	}

	return true
}

// OrGate.Exec: Returns false if no inputs are true
func (g *OrGate) Exec(circuit *Circuit) bool {
	inputs := circuit.GetInputs(g.Label)

	for _, input := range inputs {
		if input {
			return true
		}
	}

	return false
}

// NotGate.Exec: Negate the input.
func (g *NotGate) Exec(circuit *Circuit) bool {
	inputs := circuit.GetInputs(g.Label)

	if len(inputs) != 1 {
		panic("NotGate expects exactly one input")
	}

	return !inputs[0]
}

// NandGate.Exec: NAND is true if not all inputs are true
func (g *NandGate) Exec(circuit *Circuit) bool {
	inputs := circuit.GetInputs(g.Label)

	for _, input := range inputs {
		if !input {
			return true
		}
	}

	return false
}

// NorGate.Exec:  If all inputs are false, NOR is true
func (g *NorGate) Exec(circuit *Circuit) bool {
	inputs := circuit.GetInputs(g.Label)

	for _, input := range inputs {
		if input {
			return false
		}
	}

	return true
}

// XnorGate.Exec: XNOR is true if the number of true inputs is even
func (g *XnorGate) Exec(circuit *Circuit) bool {
	inputs := circuit.GetInputs(g.Label)
	trueCount := 0

	for _, input := range inputs {
		if input {
			trueCount++
		}
	}

	return trueCount%2 == 0
}

// EquGate.Exec: output the input.
func (g *EquGate) Exec(circuit *Circuit) bool {
	inputs := circuit.GetInputs(g.Label)

	if len(inputs) != 1 {
		panic("EquGate expects exactly one input")
	}

	return inputs[0]
}
