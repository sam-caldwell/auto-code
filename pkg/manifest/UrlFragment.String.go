package manifest

// String - return a string representation of an RFC-3986 Url Fragment section
func (n *UrlFragment) String() string {
	return string(*n)
}
