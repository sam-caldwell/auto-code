package manifest

// PostgresTableConstraint - represents a standardized postgresql constraint
type PostgresTableConstraint struct {
	Name       PostgresIdentifier `json:"name"`
	Nullable   bool               `json:"nullable"`
	Check      PostgresTableCheck `yaml:"check,omitempty"`
	References PostgresReference  `yaml:"references,omitempty"`
	Deferrable PostgresDeferrable `yaml:"deferrable,omitempty"`
}

// PostgresTableCheck - represents a postgresql constraint check
type PostgresTableCheck struct {
	Expression string `yaml:"expression"`
	NoInherit  bool   `yaml:"noInherit"`
}

// PostgresConstraintName - represents a constraint name
type PostgresConstraintName string

// PostgresReference - a foreign key reference
type PostgresReference struct {
	Name PostgresIdentifier
}

// PostgresDeferrable - an enumerated type representing whether a constraint is deferrable
type PostgresDeferrable int
