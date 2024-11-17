package manifest

// FromString - convert a string into its UrlFragment representation
func (n *UrlFragment) FromString(s string) {
	*n = UrlFragment(s)
}
