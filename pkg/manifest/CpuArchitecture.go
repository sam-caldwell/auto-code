package manifest

type CpuArchitecture int

const (
	_ CpuArchitecture = iota
	arm
	arm64
	amd64
)
