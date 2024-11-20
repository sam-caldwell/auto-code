package url

// FromString - capture the given string as the query part of an RFC-3986 URL
func (n *Query) FromString(s string) {
	*n = Query(s)
}
