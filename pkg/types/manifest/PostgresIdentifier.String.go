package manifest

// String - return a string representation of the standardized PostgresIdentifier
func (p *PostgresIdentifier) String() string {
	return string(*p)
}
