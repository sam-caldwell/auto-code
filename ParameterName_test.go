package arguments

import (
	"testing"
)

func TestParameterName(t *testing.T) {
	_ = ParameterName("foo") //test parameter name (but don't do validation).
	// Disclaimer: NewParameterName() should have been used in real use cases.
	//             But we want to validate type conversion from a string.
}
