package cpu

// String - return the string value of the cpu architecture
func (cpu *CpuArchitecture) String() string {
	return [...]string{"unknown", "arm", "arm64", "amd64"}[*cpu]
}
