package manifest

// FromString - capture the given string as the query part of an RFC-3986 URL
func (n *UrlQuery) FromString(s string) {
	*n = UrlQuery(s)
}
