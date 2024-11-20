package configuration

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// FileSchema - the dataContract schema for a configuration file
type FileSchema struct {
	manifest.DataObjectWithReference
	Path      generic.StructuredPropertyName `yaml:"path"`
	Parameter generic.Identifier             `yaml:"parameter"`
}
