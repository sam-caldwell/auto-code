package url

// String - return a string representation of an RFC-3986 Url Fragment section
func (n *Fragment) String() string {
	return string(*n)
}
