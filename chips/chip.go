package chips

import (
	"adder/core"

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

func (bc *Chip) Create() *core.Circuit {
	circuit := core.NewCircuit()
	for _, gate := range bc.Gates {
		circuit.AddGate(core.NewGate(gate.Type, gate.Name))
	}
	for _, conn := range bc.Connections {
		source := conn[0]
		inputs := conn[1:]
		circuit.Connect(source, inputs...)
	}
	return circuit
}

func (bc *Chip) Run(circuit *core.Circuit, inputs Ports) Ports {
	circuit.SetInputs(inputs)
	outputs := make(Ports)
	for _, output := range bc.OutPorts {
		outputs[output] = circuit.Run(output)
	}
	return outputs
}

func (bc *Chip) Write(inputs, outputs Ports) {
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
