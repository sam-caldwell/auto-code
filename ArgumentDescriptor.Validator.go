package arguments

import "fmt"

func (arg ArgumentDescriptor[I, U, F, S, B]) Validator() error {
	if arg.bounds.Validator(arg.value) {
		return nil
	}
	return fmt.Errorf(boundsCheckFailed, arg.flag, arg.value)
}
