package manifest

// String - return a string representation of a postgres table column type
func (p *PostgresTableColumnType) String() string {
	return string(*p)
}
