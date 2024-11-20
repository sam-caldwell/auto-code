package url

// FromString - convert a string into its Fragment representation
func (n *Fragment) FromString(s string) {
	*n = Fragment(s)
}
