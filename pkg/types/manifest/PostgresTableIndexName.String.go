package manifest

// String - return a string representation of a postgres table index name
func (p *PostgresTableIndexName) String() string {
	return string(*p)
}
