package arguments

import (
	"fmt"
	"testing"
)

func TestParser_Copyright(t *testing.T) {
	var expectedError = fmt.Errorf("invalid copyright string pattern")

	tests := []struct {
		input    string
		expected string
		err      error
	}{
		// Happy paths
		{"© 2023 Example Corp", "© 2023 Example Corp", nil},
		{"(c) 2023", "(c) 2023", nil},
		{"© 2022, 2023 Example Corp", "© 2022, 2023 Example Corp", nil},
		{"2023 Example Corp", "2023 Example Corp", nil},

		// Sad paths
		{"(c) 1899 Example Corp", "", expectedError}, // Year out of range
		{"(c) 2023!", "", expectedError},             // Invalid character
		{"© 20X3 Example Corp", "", expectedError},   // Invalid year format
		{"Example Corp", "", expectedError},          // Missing year
		{"", "", expectedError},                      // Empty string
	}

	for _, test := range tests {
		parser := &Parser{}
		parser = parser.Copyright(test.input)

		if parser.programCopyright != test.expected {
			t.Fatalf("expected programCopyright = %s, got %s", test.expected, parser.programCopyright)
		}

		if test.err == nil && parser.err == nil {
			continue // get the next test. we are good here.
		}
		// either test.err != nil or parser.err != nil
		if test.err == nil {
			if parser.err != nil {
				t.Fatalf("parser.Copyright(%q): expected no error %s, got an error %v",
					test.input, expectedError, parser.err)
			}
		}
		// test.err != nil
		if parser.err == nil {
			t.Fatalf("parser.Copyright(%q): expected error %s, got nil", test.input, expectedError)
		}
		// parser.err != nil
		if test.err.Error() != test.err.Error() {
			t.Fatalf("parser.Copyright(%q): expected error message(%s) but got %s",
				test.input, test.err.Error(), parser.err.Error())
		}
	}
}
