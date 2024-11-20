package docker

// String - return the string representation of the Image
func (ci *Image) String() string {
	return string(*ci)
}
