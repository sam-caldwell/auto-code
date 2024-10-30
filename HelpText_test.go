package arguments

import (
	"testing"
)

func TestHelpText(t *testing.T) {
	_ = HelpText("test string!") //this skips validation but tests that we can convert a string
}
