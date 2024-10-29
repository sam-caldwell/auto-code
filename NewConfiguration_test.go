package arguments

import (
	"testing"
)

func TestNewConfiguration(t *testing.T) {
	config := NewConfiguration()
	if config == nil {
		t.Fatal("expect non nil configuration")
	}
	if config.data == nil {
		t.Fatal("expect non nil data")
	}
}
