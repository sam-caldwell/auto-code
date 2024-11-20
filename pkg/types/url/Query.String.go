package url

// String - return the query part of an RFC-3986 URL
func (n *Query) String() string {
	return string(*n)
}
