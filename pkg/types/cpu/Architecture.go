package cpu

// Architecture - an enumerated type representing a CPU Architecture
type Architecture int

const (
	_ Architecture = iota
	arm
	arm64
	amd64
)
