package arguments

import (
	"testing"
)

func TestParser_ProgramName(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		p := Parser{}
		_ = p.ProgramName("test")
		if p.err != nil {
			t.Fatal("parser error state should be nil after ProgramName() is called")
		}
	})

	t.Run("sad path: empty name", func(t *testing.T) {
		p := Parser{}
		_ = p.ProgramName("")
		if p.err == nil {
			t.Fatal("parser error state should not be nil after ProgramName() is called with empty string")
		}
		if p.err.Error() != invalidProgramName {
			t.Fatalf("expected invalidProgramName (%s) but got '%s'", invalidProgramName, p.err.Error())
		}
	})

	t.Run("sad path: whitespace name", func(t *testing.T) {
		p := Parser{}
		_ = p.ProgramName(" ")
		if p.err == nil {
			t.Fatal("parser error state should not be nil after ProgramName() is called with whitespace")
		}
		if p.err.Error() != invalidProgramName {
			t.Fatalf("expected invalidProgramName (%s) but got '%s'", invalidProgramName, p.err.Error())
		}
	})

}
