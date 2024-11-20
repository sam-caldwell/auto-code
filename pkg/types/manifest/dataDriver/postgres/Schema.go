package postgres

import (
	"github.com/sam-caldwell/auto-code/pkg/types/manifest"
)

// DataSchemaPostgres - An implementation of DataSchemaDescriptor used for Postgresql Db
type DataSchemaPostgres struct {
	manifest.DataSchemaDescriptorBase

	Tables     []Table     `yaml:"tables,omitempty"`
	Views      []View      `yaml:"views,omitempty"`
	Functions  []Function  `yaml:"functions,omitempty"`
	Procedures []Procedure `yaml:"procedures,omitempty"`
	Enum       []Enum      `yaml:"enum,omitempty"`
}
