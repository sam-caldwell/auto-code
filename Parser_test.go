package arguments

import (
	"fmt"
	"testing"
)

func TestParser_Struct(t *testing.T) {
	_ = Parser{
		programName:        "test name",
		programDescription: "test description",
		programCopyright:   "test copyright",
		parameters:         make(map[ParameterName]ParameterList),
		err:                fmt.Errorf("test"),
	}
}
