package arguments

import (
	"testing"
)

func TestDetectArguments(t *testing.T) {
	testData := map[string]bool{
		"bad":     false,
		"-b":      true,
		"-better": true,
		"--good":  true,
		"-":       false,
		"--":      false,
		"a-":      false,
		"a--":     false,
		"a-a":     false,
		"a--a":    false,
	}
	for arg, expected := range testData {
		actual := detectArgument(&arg)
		if actual != expected {
			t.Fatalf("fail (%s): expected %v, actual %v", arg, expected, actual)
		}
	}
}
