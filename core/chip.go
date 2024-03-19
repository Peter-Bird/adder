package core

import (
	"fmt"
)

type Gates []struct {
	Type string
	Name string
}

type Connections [][]string

type PortList []string

type Chip struct {
	GateType    string
	Gates       Gates
	Connections Connections
	InPorts     PortList
	OutPorts    PortList
}

type Ports map[string]bool

func (c *Chip) Create() *Circuit {
	circuit := NewCircuit()
	for _, gate := range c.Gates {
		circuit.AddGate(NewGate(gate.Type, gate.Name))
	}
	for _, conn := range c.Connections {
		source := conn[0]
		inputs := conn[1:]
		circuit.Connect(source, inputs...)
	}
	return circuit
}

func (c *Chip) Run(circuit *Circuit, inputs Ports) Ports {
	circuit.SetInputs(inputs)
	outputs := make(Ports)
	for _, output := range c.OutPorts {
		outputs[output] = circuit.Run(output)
	}
	return outputs
}

func (c *Chip) Write(inputs, outputs Ports) {
	// Print inputs
	fmt.Println("In:")
	for key, value := range inputs {
		fmt.Printf("\t%s: %t\n", key, value)
	}

	// Print outputs
	fmt.Println("Out:")
	for key, value := range outputs {
		fmt.Printf("\t%s: %t\n", key, value)
	}
}

func printPorts(name string, ports Ports) {

	fmt.Println(name)
	for key, value := range ports {
		fmt.Printf("\t%s: %t\n", key, value)
	}
}
