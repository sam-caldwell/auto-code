package arguments

import (
	"fmt"
)

func (arg *ArgumentDescriptor[I, U, F, S, B]) String() string {
	return fmt.Sprintf("%v (default: '%v'): %v", arg.flag, arg.value, arg.help)
}
