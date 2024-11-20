package artifact

import (
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// Descriptor - represents a single project artifact.
//
// An artifact may be explicitly defined or may reference an external document
// which will be consumed to populate the data.
type Descriptor struct {
	manifest.DataObjectWithReference
	Name         Name                   `yaml:"name,omitempty"`
	Description  generic.NonEmptyString `yaml:"description,omitempty"`
	Type         Type                   `yaml:"type,omitempty"`
	Dependencies []Name                 `yaml:"dependencies,omitempty"`
	Options      []Options              `yaml:"options,omitempty"`
}

// Generate - Generate the project artifact
func (a *Descriptor) Generate(target string, debug bool) error {
	if debug {
		ansi.Blue().Printf("Generate a %s: %s", a.Type.String(), a.Name.String())
	}
	switch a.Type {
	case Service:
	case External:
	case Binary:
	default:
		ansi.Red().Printf("unrecognized artifact type:'%s'" + a.Type.String()).LF().Fatal(1)
	}
	return nil
}
