package arguments

import (
	"testing"
)

func TestNewParser(t *testing.T) {
	p := NewParser()
	if p == nil {
		t.Fatal("expected non-nil parser")
	}
	if p.err != nil {
		t.Fatal("expected nil parser error")
	}
	if p.parameters == nil {
		t.Fatal("expected non-nil args")
	}
	t.Log("this test needs close inspection")
}
