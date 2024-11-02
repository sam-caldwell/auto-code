package validator

import (
	"math"
	"testing"
)

func TestInt(test *testing.T) {
	_ = Int{
		Min: math.MinInt,
		Max: math.MaxInt,
	}
}
