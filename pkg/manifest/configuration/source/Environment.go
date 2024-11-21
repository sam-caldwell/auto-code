package source

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// Environment - represents an environment variable for the Configuration object
type Environment struct {
	manifest.DataObjectWithReference
	Path      generic.Identifier `yaml:"path"`
	Parameter generic.Identifier `yaml:"parameter"`
}
