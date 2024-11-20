package postgres

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
)

// Table - represents a database table
type Table struct {
	Name        dataCommon.Identifier   `yaml:"name"`
	Description manifest.NonEmptyString `yaml:"description,omitempty"`
	Columns     []TableColumn           `yaml:"columns"`
	Indexes     []TableIndex            `yaml:"indexes,omitempty"`
}
