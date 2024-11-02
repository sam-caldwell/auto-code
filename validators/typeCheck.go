package validator

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"golang.org/x/exp/constraints"
	"reflect"
)

func minMaxTypeCheck[T constraints.Ordered](min T, max T, expectedType string) error {

	if reflect.TypeOf(min).String() != expectedType {
		return fmt.Errorf(messages.ErrTypeMismatchMin)
	}

	if reflect.TypeOf(max).String() != expectedType {
		return fmt.Errorf(messages.ErrTypeMismatchMax)
	}

	if max <= min {
		return fmt.Errorf(messages.ErrInvalidMinMaxValue)
	}
	return nil
}
