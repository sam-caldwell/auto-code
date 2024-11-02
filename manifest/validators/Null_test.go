package validator

import (
	"reflect"
	"testing"
)

func TestNull(t *testing.T) {
	t.Run("basic validation", func(t *testing.T) {
		_ = Null{}
	})
	t.Run("verify there are no fields in the Null struct", func(t *testing.T) {
		oType := reflect.TypeOf(Null{})
		if numFields := oType.NumField(); numFields != 0 {
			t.Fatalf("expected 0 fields, but found %d", numFields)
		}
	})
}
