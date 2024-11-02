package validator

import (
	"fmt"
	"reflect"
)

// Verify - Verify the Int16 YAML Manifest Validator Parameter
//
// Things to verify...
// - Make sure min < max
// - verify the parameter is big enough to contain its data range.
func (p Int16) Verify() error {
	if reflect.TypeOf(p.Min).String() != "int16" {
		return fmt.Errorf("type mismatch (Min)")
	}
	if reflect.TypeOf(p.Max).String() != "int16" {
		return fmt.Errorf("type mismatch (Max)")
	}
	if p.Max <= p.Min {
		return fmt.Errorf("min must be less than max")
	}
	return nil
}
