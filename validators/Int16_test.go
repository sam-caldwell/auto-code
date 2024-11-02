package validator

import (
	"math"
	"testing"
)

func TestInt16(test *testing.T) {
	_ = Int16{
		Min: math.MinInt16,
		Max: math.MaxInt16,
	}
}
