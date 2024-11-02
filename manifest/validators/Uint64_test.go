package validator

import (
	"math"
	"testing"
)

func TestUint64(test *testing.T) {
	_ = Uint64{
		Min: 0,
		Max: math.MaxUint64,
	}
}
