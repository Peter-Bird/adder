package chips

import (
	"adder/core"

	"fmt"
)

// Chip defines a standard interface for chip operations.
type Chip interface {
	Create() *core.Circuit
	Run(*core.Circuit, map[string]bool) interface{}
	Template() string
	Write(map[string]bool, interface{})
}

type Gates []struct {
	Type string
	Name string
}

type Connections [][]string

type Ports []string

func Write(inputs map[string]bool, outputs map[string]bool) {
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
