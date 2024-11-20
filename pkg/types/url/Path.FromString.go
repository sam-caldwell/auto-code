package url

// FromString - cast a string into a Path (RFC 3986)
func (p *Path) FromString(pathString string) {
	//ToDo: validate this.
	*p = Path(pathString)
}
