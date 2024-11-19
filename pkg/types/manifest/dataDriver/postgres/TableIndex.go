package postgres

// TableIndex - represents a standardized postgresql index
type TableIndex struct {
	Name Identifier `yaml:"name"`
}
