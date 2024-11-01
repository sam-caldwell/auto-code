package manifest

import "fmt"

// Validate - Verify that min < max
func (p ParameterMinMax[T]) Validate(propertyType *string) (err error) {
	if p.Min < p.Max {
		return nil
	}
	return fmt.Errorf(errInvalidMinMaxValue)
}
