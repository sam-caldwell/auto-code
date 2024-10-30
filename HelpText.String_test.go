package arguments

import (
	"testing"
)

func TestHelpText_String(t *testing.T) {
	ht := HelpText("test string")
	if ht.String() != "test string" {
		t.Fatalf("expected input not recd")
	}
}
