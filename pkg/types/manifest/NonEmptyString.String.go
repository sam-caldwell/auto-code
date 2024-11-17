package manifest

// String returns the string representation of the NonEmptyString
func (n *NonEmptyString) String() string {
	return string(*n)
}
