package manifest

import (
	"testing"
)

func TestCommandlineArgList(t *testing.T) {
	_ = CommandlineArgList{
		"foobar": Commandline{
			Required: true,
			Short:    "-f",
			Long:     "--foobar",
		},
	}
}
