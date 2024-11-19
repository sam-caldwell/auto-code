package postgres

// String - return a string representation of a postgres data type
func (p *Type) String() string {
	return string(*p)
}
