package network

// String - return string representation of network address part of RFC-3986 URL
func (n *Address) String() string {
	return string(*n)
}
