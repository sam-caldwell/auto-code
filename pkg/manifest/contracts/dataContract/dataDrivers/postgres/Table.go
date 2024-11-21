package postgres

import (
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// Table - represents a database table
type Table struct {
	Name        dataCommon.Identifier  `yaml:"name"`
	Description generic.NonEmptyString `yaml:"description,omitempty"`
	Columns     []TableColumn          `yaml:"columns"`
	Indexes     []TableIndex           `yaml:"indexes,omitempty"`
}
