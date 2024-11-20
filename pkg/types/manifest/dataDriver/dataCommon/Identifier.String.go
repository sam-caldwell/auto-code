package dataDriver

// String - return a string representation of the standardized PostgresIdentifier
func (p *Identifier) String() string {
	return string(*p)
}
