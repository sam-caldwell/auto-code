package validator

import "testing"

func TestNull_Verify(t *testing.T) {
	o := Null{}
	if err := o.Verify(); err != nil {
		t.Error(err)
	}
}
