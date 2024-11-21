package source

import (
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// CommandLineArg - represents a command-line argument for the Configuration object
//
// This object represents an argument (Name) and associated Configuration Parameter.
// When processed, this will use the Name to generate the long-form argument (--${Name})
// for any command-line args.
//
// Note: Parameter must exist in the Configuration.Parameters section.
type CommandLineArg struct {
	xref.DataObjectWithReference
	Name      generic.Identifier `yaml:"Name"`
	Parameter generic.Identifier `yaml:"parameter"`
}
