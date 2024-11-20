package postgres

// String - return a string representation of a postgres dataContract type
func (p *Type) String() string {
	return string(*p)
}
