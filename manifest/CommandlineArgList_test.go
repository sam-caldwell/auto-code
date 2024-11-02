package manifest

import (
	"testing"
)

func TestCommandlineArgList(t *testing.T) {
	arg := CommandlineArgList{}
	arg = append(arg, Commandline{})
}
