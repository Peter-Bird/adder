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

func Fmt(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}
