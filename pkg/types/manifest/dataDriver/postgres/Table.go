package postgres

import "github.com/sam-caldwell/auto-code/pkg/types/manifest"

// Table - represents a database table
type Table struct {
	Name        Identifier              `yaml:"name"`
	Description manifest.NonEmptyString `yaml:"description,omitempty"`
	Columns     []TableColumn           `yaml:"columns"`
	Indexes     []TableIndex            `yaml:"indexes,omitempty"`
}
