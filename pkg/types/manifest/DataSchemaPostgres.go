package manifest

// DataSchemaPostgres - An implementation of DataSchemaDescriptor used for Postgresql Db
type DataSchemaPostgres struct {
	DataSchemaDescriptorBase

	Tables     []PostgresTable     `yaml:"tables,omitempty"`
	Views      []PostgresView      `yaml:"views,omitempty"`
	Functions  []PostgresFunction  `yaml:"functions,omitempty"`
	Procedures []PostgresProcedure `yaml:"procedures,omitempty"`
	Enum       []PostgresEnum      `yaml:"enum,omitempty"`
}
