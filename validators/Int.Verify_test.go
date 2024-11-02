package validator

import (
	"testing"
)

func TestInt_Verify(t *testing.T) {

	t.Run("happy: min < max", func(t *testing.T) {
		o := Int{
			Min: 0,
			Max: 10,
		}
		if err := o.Verify(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("sad: min==max", func(t *testing.T) {
		o := Int{
			Min: 10,
			Max: 10,
		}
		if err := o.Verify(); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("sad: max<min", func(t *testing.T) {
		o := Int{
			Min: 10,
			Max: 0,
		}
		if err := o.Verify(); err == nil {
			t.Fatal("expected error")
		}
	})
}
