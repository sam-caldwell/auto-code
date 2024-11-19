package manifest

// PostgresTable - represents a database table
type PostgresTable struct {
	Name        PostgresIdentifier    `yaml:"name"`
	Description NonEmptyString        `yaml:"description,omitempty"`
	Columns     []PostgresTableColumn `yaml:"columns"`
	Indexes     []PostgresTableIndex  `yaml:"indexes,omitempty"`
}
