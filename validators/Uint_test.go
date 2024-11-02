package validator

import (
	"math"
	"testing"
)

func TestUint(test *testing.T) {
	_ = Uint{
		Min: 0,
		Max: math.MaxUint,
	}
}
