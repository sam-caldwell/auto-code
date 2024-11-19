package manifest

// PostgresTableIndex - represents a standardized postgresql index
type PostgresTableIndex struct {
	Name PostgresIdentifier `yaml:"name"`
}
