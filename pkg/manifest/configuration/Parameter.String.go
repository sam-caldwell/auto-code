package configuration

// String - return a string representation of a config parameter
func (p *Parameter) String() string {
	return p.Value.String()
}
