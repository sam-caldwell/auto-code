package manifest

// String - return the string representation of the ContainerImage
func (ci *ContainerImage) String() string {
	return string(*ci)
}
