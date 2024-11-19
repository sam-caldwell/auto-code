package manifest

// PostgresTableColumn - represents a standardized postgresql table column
type PostgresTableColumn struct {
	Name        PostgresIdentifier        `yaml:"name"`
	Type        PostgresType              `yaml:"type"`
	Nullable    bool                      `yaml:"nullable,omitempty"`
	Constraints []PostgresTableConstraint `yaml:"constraints,omitempty"`
}
