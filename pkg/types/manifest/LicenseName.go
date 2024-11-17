package manifest

// LicenseName - represents an enumerated License type
type LicenseName int

const (
	Proprietary LicenseName = iota
	MIT
	Apache2
	BSD
)
