package validator

import (
	"fmt"
	"reflect"
)

// Verify - Verify the Uint YAML Manifest Validator Parameter
//
// Things to verify...
// - Make sure min < max
// - verify the parameter is big enough to contain its data range.
func (p Uint) Verify() error {
	if typ := reflect.TypeOf(p.Min).String(); typ != "uint" {
		return fmt.Errorf("type mismatch (Min) got: %s", typ)
	}
	if typ := reflect.TypeOf(p.Max).String(); typ != "uint" {
		return fmt.Errorf("type mismatch (Max) got: %s", typ)
	}
	if p.Max <= p.Min {
		return fmt.Errorf("min must be less than max")
	}
	return nil
}
