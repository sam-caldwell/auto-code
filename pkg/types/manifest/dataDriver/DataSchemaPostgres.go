package dataDriver

import (
	"github.com/sam-caldwell/auto-code/pkg/types/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/manifest/dataDriver/postgres"
)

// DataSchemaPostgres - An implementation of DataSchemaDescriptor used for Postgresql Db
type DataSchemaPostgres struct {
	manifest.DataSchemaDescriptorBase

	Tables     []postgres.Table     `yaml:"tables,omitempty"`
	Views      []postgres.View      `yaml:"views,omitempty"`
	Functions  []postgres.Function  `yaml:"functions,omitempty"`
	Procedures []postgres.Procedure `yaml:"procedures,omitempty"`
	Enum       []postgres.Enum      `yaml:"enum,omitempty"`
}
