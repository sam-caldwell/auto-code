package source

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// CommandLineArg - represents a command-line argument for the Configuration object
type CommandLineArg struct {
	manifest.DataObjectWithReference
	Path      generic.Identifier `yaml:"path"`
	Parameter generic.Identifier `yaml:"parameter"`
}
