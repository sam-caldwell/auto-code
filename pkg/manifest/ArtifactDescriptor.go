package manifest

import (
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// ArtifactDescriptor - represents a single project artifact.
//
// An artifact may be explicitly defined or may reference an external document
// which will be consumed to populate the data.
type ArtifactDescriptor struct {
	DataObjectWithReference
	Name         ArtifactName           `yaml:"name,omitempty"`
	Description  generic.NonEmptyString `yaml:"description,omitempty"`
	Type         ArtifactType           `yaml:"type,omitempty"`
	Dependencies []ArtifactName         `yaml:"dependencies,omitempty"`
	Options      []ArtifactOptions      `yaml:"options,omitempty"`
}

// Generate - Generate the project artifact
func (a *ArtifactDescriptor) Generate(target string, debug bool) error {
	if debug {
		ansi.Blue().Printf("Generate a %s: %s", a.Type.String(), a.Name.String())
	}
	switch a.Type {
	case service:
	case external:
	case binary:
	default:
		ansi.Red().Printf("unrecognized artifact type:'%s'" + a.Type.String()).LF().Fatal(1)
	}
	return nil
}
