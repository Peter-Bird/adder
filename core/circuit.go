package core

type Circuit struct {
	Gates       map[string]Gate
	Connections map[string][]string
	Inputs      map[string]bool
}

func NewCircuit() *Circuit {
	return &Circuit{
		Gates:       make(map[string]Gate),
		Connections: make(map[string][]string),
		Inputs:      make(map[string]bool),
	}
}

func (c *Circuit) AddGate(gate Gate) {
	c.Gates[gate.GetLabel()] = gate
}

func (c *Circuit) AddGates(gates []Gate) {
	for _, gate := range gates {
		c.AddGate(gate)
	}
}

func (c *Circuit) Connect(gate string, inputs ...string) {
	c.Connections[gate] = inputs
}

func (c *Circuit) SetInput(label string, value bool) {
	c.Inputs[label] = value
}

func (c *Circuit) SetInputs(inputs map[string]bool) {
	for label, value := range inputs {
		c.Inputs[label] = value
	}
}

func (c *Circuit) GetInputs(gateLabel string) []bool {
	inputLabels := c.Connections[gateLabel]
	inputs := make([]bool, len(inputLabels))

	for i, label := range inputLabels {
		input, ok := c.Inputs[label]
		if !ok {
			// If the input is not found in the Inputs map,
			// assume it is an output from another gate
			// and execute that gate.
			input = c.Gates[label].Exec(c)
		}
		inputs[i] = input
	}

	return inputs
}

func (c *Circuit) Run(gateLabel string) bool {
	gate, exists := c.Gates[gateLabel]
	if !exists {
		return false
	}
	return gate.Exec(c)
}
