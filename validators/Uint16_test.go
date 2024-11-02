package validator

import (
	"math"
	"testing"
)

func TestUint16(test *testing.T) {
	_ = Uint16{
		Min: 0,
		Max: math.MaxUint16,
	}
}
