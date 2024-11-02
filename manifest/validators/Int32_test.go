package validator

import (
	"math"
	"testing"
)

func TestInt32(test *testing.T) {
	_ = Int32{
		Min: math.MinInt32,
		Max: math.MaxInt32,
	}
}
