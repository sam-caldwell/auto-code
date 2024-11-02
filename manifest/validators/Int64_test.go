package validator

import (
	"math"
	"testing"
)

func TestInt64(test *testing.T) {
	_ = Int64{
		Min: math.MinInt64,
		Max: math.MaxInt64,
	}
}
