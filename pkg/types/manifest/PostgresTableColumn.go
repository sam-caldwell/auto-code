package manifest

// PostgresTableColumn - represents a standardized postgresql table column
type PostgresTableColumn struct {
	Name        PostgresColumnName        `yaml:"name"`
	Type        PostgresColumnType        `yaml:"type"`
	Nullable    bool                      `yaml:"nullable,omitempty"`
	Constraints []PostgresTableConstraint `yaml:"constraints,omitempty"`
}
