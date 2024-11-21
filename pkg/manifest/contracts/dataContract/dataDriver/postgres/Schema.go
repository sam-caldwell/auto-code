package postgres

import "github.com/sam-caldwell/auto-code/pkg/types/xref"

// Schema - A PostgreSQL implementation of DataSchemaDescriptor.
type Schema struct {
	xref.DataObjectWithReference
	Tables     []Table     `yaml:"tables,omitempty"`
	Views      []View      `yaml:"views,omitempty"`
	Functions  []Function  `yaml:"functions,omitempty"`
	Procedures []Procedure `yaml:"procedures,omitempty"`
	Enum       []Enum      `yaml:"enum,omitempty"`
}
