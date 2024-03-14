package main

import (
	"fmt"
)

// Gate interface that all gates implement
type Gate interface {
	Exec(circuit *Circuit) bool
	Label() string
}

// BasicGate contains common properties,
// including the gate's label
type BasicGate struct {
	label string
}

func (g *BasicGate) Label() string {
	return g.label
}

type XorGate struct{ BasicGate }
type AndGate struct{ BasicGate }
type OrGate struct{ BasicGate }

func (g *XorGate) Exec(circuit *Circuit) bool {
	input1, input2 := circuit.GetInputs(g.label)
	return input1 != input2
}

func (g *AndGate) Exec(circuit *Circuit) bool {
	input1, input2 := circuit.GetInputs(g.label)
	return input1 && input2
}

func (g *OrGate) Exec(circuit *Circuit) bool {
	input1, input2 := circuit.GetInputs(g.label)
	return input1 || input2
}

// Circuit contains the gates,
// their connections,
// and a method to run the circuit
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

func (c *Circuit) AddGate(gate Gate) {
	c.Gates[gate.Label()] = gate
}

// Connect sets the inputs for a gate,
// identified by gate labels or
// direct input values ("true", "false")
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

func (c *Circuit) Run(gateLabel string) bool {
	return c.Gates[gateLabel].Exec(c)
}

func main() {
	// Initialize the circuit
	circuit := NewCircuit()

	// Add gates to the circuit
	circuit.AddGate(&XorGate{BasicGate{"XOR1"}})
	circuit.AddGate(&XorGate{BasicGate{"XOR2"}})
	circuit.AddGate(&AndGate{BasicGate{"AND1"}})
	circuit.AddGate(&AndGate{BasicGate{"AND2"}})
	circuit.AddGate(&OrGate{BasicGate{"OR"}})

	// Define connections (for a full adder circuit)
	circuit.SetInput("A", false)
	circuit.SetInput("B", false)
	circuit.SetInput("Cin", false)

	circuit.Connect("XOR1", "A", "B")
	circuit.Connect("XOR2", "XOR1", "Cin")
	circuit.Connect("AND1", "A", "B")
	circuit.Connect("AND2", "XOR1", "Cin")
	circuit.Connect("OR", "AND1", "AND2")

	// Run the circuit to get the sum and carry out
	sum := circuit.Run("XOR2")
	carryOut := circuit.Run("OR")

	// Print the results
	fmt.Printf("Sum: %t, Carry Out: %t\n", sum, carryOut)
}
