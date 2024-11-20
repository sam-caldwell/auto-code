package license

// Name - represents an enumerated License type
type Name int

const (
	Proprietary Name = iota
	MIT
	Apache2
	BSD
)
