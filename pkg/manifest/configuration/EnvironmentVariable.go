package configuration

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// EnvironmentVariable - represents an environment variable for the Configuration object
type EnvironmentVariable struct {
	manifest.DataObjectWithReference
	Path      generic.Identifier `yaml:"path"`
	Parameter generic.Identifier `yaml:"parameter"`
}
