package chips

import "fmt"

func Fmt(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}
