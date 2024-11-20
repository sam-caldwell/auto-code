package postgres

// TableIndex - represents a standardized postgresql index
type TableIndex struct {
	Name dataCommon.Identifier `yaml:"name"`
}
