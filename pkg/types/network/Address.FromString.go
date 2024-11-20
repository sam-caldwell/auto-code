package network

// FromString - capture the network address (host) part of an RFC-3986 URL
func (n *Address) FromString(s string) {
	*n = Address(s)
}
