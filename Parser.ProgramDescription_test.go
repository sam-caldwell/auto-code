package arguments

import (
	"testing"
)

func TestParser_ProgramDescription(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		p := Parser{}
		_ = p.ProgramDescription("test")
		if p.err != nil {
			t.Fatal("parser error state should be nil after ProgramDescription() is called")
		}
	})

	t.Run("sad path: empty name", func(t *testing.T) {
		p := Parser{}
		_ = p.ProgramDescription("")
		if p.err == nil {
			t.Fatal("parser error state should not be nil after ProgramDescription() is called with empty string")
		}
	})

	t.Run("sad path: whitespace name", func(t *testing.T) {
		p := Parser{}
		_ = p.ProgramDescription(" ")
		if p.err == nil {
			t.Fatal("parser error state should not be nil after ProgramDescription() is called with whitespace")
		}
	})

}
