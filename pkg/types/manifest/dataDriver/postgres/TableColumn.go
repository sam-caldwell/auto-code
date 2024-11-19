package postgres

// TableColumn - represents a standardized postgresql table column
type TableColumn struct {
	Name        Identifier        `yaml:"name"`
	Type        Type              `yaml:"type"`
	Nullable    bool              `yaml:"nullable,omitempty"`
	Constraints []TableConstraint `yaml:"constraints,omitempty"`
}
