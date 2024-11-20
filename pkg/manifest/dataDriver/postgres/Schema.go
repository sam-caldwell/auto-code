package postgres

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/dataContract"
)

// Schema - A PostgreSQL implementation of DataSchemaDescriptor.
type Schema struct {
	dataContract.SchemaDescriptorBase

	Tables     []Table     `yaml:"tables,omitempty"`
	Views      []View      `yaml:"views,omitempty"`
	Functions  []Function  `yaml:"functions,omitempty"`
	Procedures []Procedure `yaml:"procedures,omitempty"`
	Enum       []Enum      `yaml:"enum,omitempty"`
}
