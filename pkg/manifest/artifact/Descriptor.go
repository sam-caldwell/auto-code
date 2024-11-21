package artifact

import (
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// Descriptor - represents a single project artifact.
//
// An artifact may be explicitly defined or may reference an external document
// which will be consumed to populate the dataContract.
type Descriptor struct {
	xref.DataObjectWithReference
	Name         generic.Identifier     `yaml:"name,omitempty"`
	Description  generic.NonEmptyString `yaml:"description,omitempty"`
	Type         Type                   `yaml:"type,omitempty"`
	Dependencies []generic.Identifier   `yaml:"dependencies,omitempty"`
	Options      []Options              `yaml:"options,omitempty"`
}
