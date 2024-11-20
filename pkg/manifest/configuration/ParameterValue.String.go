package configuration

// String - return a string representation of the ParameterValue
func (p *ParameterValue) String() string {
	return p.Data.String()
}
