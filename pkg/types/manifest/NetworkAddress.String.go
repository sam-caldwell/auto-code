package manifest

// String - return string representation of network address part of RFC-3986 URL
func (n *NetworkAddress) String() string {
	return string(*n)
}
