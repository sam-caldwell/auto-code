package source

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// FileSchema - define the configuration file parameter set.
//
// Each object is an identifier (property path) in a configuration file, which
// defines a specific configuration property the user's project manifest can set.
type FileSchema struct {
	manifest.DataObjectWithReference
	Name      generic.StructuredPropertyName `yaml:"name"`
	Parameter generic.Identifier             `yaml:"parameter"`
}
