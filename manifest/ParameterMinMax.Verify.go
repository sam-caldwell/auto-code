package manifest

import "fmt"

// Verify - validate manifest.yaml parameter to ensure that min < max
func (p ParameterMinMax[T]) Verify() (err error) {
	if p.Min < p.Max {
		return nil
	}
	return fmt.Errorf(errInvalidMinMaxValue)
}
