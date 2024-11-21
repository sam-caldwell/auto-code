package source

import (
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// Environment - represents an environment variable for the Configuration object
type Environment struct {
	xref.DataObjectWithReference
	Name      generic.Identifier `yaml:"name"`
	Parameter generic.Identifier `yaml:"parameter"`
}
