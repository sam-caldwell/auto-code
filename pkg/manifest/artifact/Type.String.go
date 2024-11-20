package artifact

// String - return a string representation of the Type
func (t *Type) String() string {

	return [...]string{
		"service",
		"external",
		"binary",
	}[int(*t)]
}
