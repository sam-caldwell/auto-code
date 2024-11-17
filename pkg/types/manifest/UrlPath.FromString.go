package manifest

// FromString - cast a string into a UrlPath (RFC 3986)
func (p *UrlPath) FromString(pathString string) {
	//ToDo: validate this.
	*p = UrlPath(pathString)
}
