package manifest

// String - return the query part of an RFC-3986 URL
func (n *UrlQuery) String() string {
	return string(*n)
}
