package validator

import (
	"fmt"
	"reflect"
)

// Verify - Verify the Uint16 YAML Manifest Validator Parameter
//
// Things to verify...
// - Make sure min < max
// - verify the parameter is big enough to contain its data range.
func (p Uint16) Verify() error {
	if reflect.TypeOf(p.Min).String() != "uint16" {
		return fmt.Errorf("type mismatch (Min)")
	}
	if reflect.TypeOf(p.Max).String() != "uint16" {
		return fmt.Errorf("type mismatch (Max)")
	}
	if p.Max <= p.Min {
		return fmt.Errorf("min must be less than max")
	}
	return nil
}
