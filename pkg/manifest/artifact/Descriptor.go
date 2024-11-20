package artifact

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// Descriptor - represents a single project artifact.
//
// An artifact may be explicitly defined or may reference an external document
// which will be consumed to populate the dataContract.
type Descriptor struct {
	manifest.DataObjectWithReference
	Name         generic.Identifier     `yaml:"name,omitempty"`
	Description  generic.NonEmptyString `yaml:"description,omitempty"`
	Type         Type                   `yaml:"type,omitempty"`
	Dependencies []generic.Identifier   `yaml:"dependencies,omitempty"`
	Options      []Options              `yaml:"options,omitempty"`
}
