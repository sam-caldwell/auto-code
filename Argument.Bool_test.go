package arguments

import (
	"testing"
)

func Test_Argument_Bool(t *testing.T) {
	arg := NewParser().
		Bool("--bool", true, "test boolean")
	if arg.flag != "--bool" {

	}
}
