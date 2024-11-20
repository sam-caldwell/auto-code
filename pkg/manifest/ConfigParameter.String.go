package manifest

// String - return a string representation of a config parameter
func (p *ConfigParameter) String() string {
	return p.Value.String()
}
