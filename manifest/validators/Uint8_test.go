package validator

import (
	"math"
	"testing"
)

func TestUint8(test *testing.T) {
	_ = Uint8{
		Min: 0,
		Max: math.MaxUint8,
	}
}
