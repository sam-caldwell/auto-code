package postgres

import (
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// TableConstraint - represents a standardized postgresql constraint
type TableConstraint struct {
	Name       generic.Identifier `json:"name"`
	Nullable   bool               `json:"nullable"`
	Check      TableCheck         `yaml:"check,omitempty"`
	References Reference          `yaml:"references,omitempty"`
	Deferrable Deferrable         `yaml:"deferrable,omitempty"`
}

// TableCheck - represents a postgresql constraint check
type TableCheck struct {
	Expression string `yaml:"expression"`
	NoInherit  bool   `yaml:"noInherit"`
}

// ConstraintName - represents a constraint name
type ConstraintName string

// Reference - a foreign key reference
type Reference struct {
	Name generic.Identifier
}

// Deferrable - an enumerated type representing whether a constraint is deferrable
type Deferrable int
