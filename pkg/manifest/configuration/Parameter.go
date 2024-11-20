package configuration

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// Parameter - represents the top-level object in Configuration.
// The ConfigParameters type represents the dataContract object used by the application
// to reference fully loaded, resolved and validated configuration parameters.
type Parameter struct {
	manifest.DataObjectWithReference
	Name        generic.Identifier      `yaml:"name"`
	Description generic.NonEmptyString  `yaml:"description"`
	Value       manifest.ParameterValue `yaml:"value"`
}
