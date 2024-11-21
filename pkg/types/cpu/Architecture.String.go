package cpu

// String - return the string value of the cpu architecture
func (cpu *Architecture) String() string {
	return [...]string{"unknown", "arm", "arm64", "amd64"}[*cpu]
}
