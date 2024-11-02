package validator

import (
	"math"
	"testing"
)

func TestInt8(test *testing.T) {
	_ = Int8{
		Min: math.MinInt8,
		Max: math.MaxInt8,
	}
}
