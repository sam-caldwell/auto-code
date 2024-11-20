package manifest

// FromString - capture the network address (host) part of an RFC-3986 URL
func (n *NetworkAddress) FromString(s string) {
	*n = NetworkAddress(s)
}
