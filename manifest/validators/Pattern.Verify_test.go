package validator

import (
	"testing"
)

func TestPattern_Verify(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		p := Pattern{
			Regex: "[a-z]+",
		}
		if err := p.Verify(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("sad: bad regex", func(t *testing.T) {
		p := Pattern{
			Regex: "][a-z][+",
		}
		if err := p.Verify(); err == nil {
			t.Fatal("expected error")
		}
	})
	t.Run("sad: empty regex", func(t *testing.T) {
		p := Pattern{
			Regex: "",
		}
		if err := p.Verify(); err == nil {
			t.Fatal("expected error")
		}
	})
}
