package core

type Gate interface {
	Exec(circuit *Circuit) bool
	GetLabel() string
}

func NewGate(gateType, label string) Gate {
	switch gateType {
	case "XorGate":
		return &XorGate{BasicGate: BasicGate{Label: label}}
	case "AndGate":
		return &AndGate{BasicGate: BasicGate{Label: label}}
	case "OrGate":
		return &OrGate{BasicGate: BasicGate{Label: label}}
	case "NotGate":
		return &NotGate{BasicGate: BasicGate{Label: label}}
	case "NandGate":
		return &NandGate{BasicGate: BasicGate{Label: label}}
	case "NorGate":
		return &NorGate{BasicGate: BasicGate{Label: label}}
	case "XnorGate":
		return &XnorGate{BasicGate: BasicGate{Label: label}}
	default:
		return nil
	}
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

func (g *XorGate) Exec(circuit *Circuit) bool {
	input1, input2 := circuit.GetInputs(g.Label)
	return input1 != input2
}

func (g *AndGate) Exec(circuit *Circuit) bool {
	input1, input2 := circuit.GetInputs(g.Label)
	return input1 && input2
}

func (g *OrGate) Exec(circuit *Circuit) bool {
	input1, input2 := circuit.GetInputs(g.Label)
	return input1 || input2
}

func (g *NotGate) Exec(circuit *Circuit) bool {
	input := circuit.GetInput(g.Label)
	return !input
}

func (g *NandGate) Exec(circuit *Circuit) bool {
	input1, input2 := circuit.GetInputs(g.Label)
	return !(input1 && input2)
}

func (g *NorGate) Exec(circuit *Circuit) bool {
	input1, input2 := circuit.GetInputs(g.Label)
	return !(input1 || input2)
}

func (g *XnorGate) Exec(circuit *Circuit) bool {
	input1, input2 := circuit.GetInputs(g.Label)
	return input1 == input2
}
