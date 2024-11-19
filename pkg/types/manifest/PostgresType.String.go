package manifest

// String - return a string representation of a postgres data type
func (p *PostgresType) String() string {
	return string(*p)
}
