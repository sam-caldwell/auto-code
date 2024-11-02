package validator

import (
	"math"
	"testing"
)

func TestUint32(test *testing.T) {
	_ = Uint32{
		Min: 0,
		Max: math.MaxUint32,
	}
}
