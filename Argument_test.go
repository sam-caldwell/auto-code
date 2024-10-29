package arguments

import (
	"fmt"
	"testing"
)

func TestArgumentStruct(t *testing.T) {
	_ = Argument{
		err:         fmt.Errorf("test"),
		programName: "test",
		copyright:   "test",
		args:        make(map[string]*ArgumentDescriptor[int64, uint64, float64, string, bool]),
	}
}
