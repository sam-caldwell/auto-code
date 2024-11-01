package manifest

import (
	"golang.org/x/exp/constraints"
)

// ParameterMinMax - defines the min and max parameters for validation.
type ParameterMinMax[T constraints.Ordered] struct {
	Min T `yaml:"min"`
	Max T `yaml:"max"`
}
