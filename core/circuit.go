package core

type Circuit struct {
	Gates       map[string]Gate
	Connections map[string][2]string
	Inputs      map[string]bool
}

func NewCircuit() *Circuit {
	return &Circuit{
		Gates:       make(map[string]Gate),
		Connections: make(map[string][2]string),
		Inputs:      make(map[string]bool),
	}
}

func (c *Circuit) AddGates(gates []Gate) {
	for _, gate := range gates {
		c.AddGate(gate)
	}
}

func (c *Circuit) AddGate(gate Gate) {
	c.Gates[gate.GetLabel()] = gate
}

func (c *Circuit) Connect(gate string, in1, in2 string) {
	c.Connections[gate] = [2]string{in1, in2}
}

func (c *Circuit) SetInput(label string, value bool) {
	c.Inputs[label] = value
}

func (c *Circuit) GetInputs(gateLabel string) (bool, bool) {
	inputLabels := c.Connections[gateLabel]
	input1, ok := c.Inputs[inputLabels[0]]
	if !ok {
		input1 = c.Gates[inputLabels[0]].Exec(c)
	}
	input2, ok := c.Inputs[inputLabels[1]]
	if !ok {
		input2 = c.Gates[inputLabels[1]].Exec(c)
	}
	return input1, input2
}

func (c *Circuit) GetInput(gateLabel string) bool {
	inputLabels := c.Connections[gateLabel]

	input, ok := c.Inputs[inputLabels[0]]
	if !ok {
		input = c.Gates[inputLabels[0]].Exec(c)
	}
	return input
}

func (c *Circuit) SetInputs(inputs map[string]bool) {
	for label, value := range inputs {
		c.SetInput(label, value)
	}
}

func (c *Circuit) Run(gateLabel string) bool {
	return c.Gates[gateLabel].Exec(c)
}
