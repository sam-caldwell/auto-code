package manifest

import "github.com/sam-caldwell/auto-code/pkg/types/generic"

// ConfigParameter - represents the top-level object in Configuration.
// The ConfigParameters type represents the data object used by the application
// to reference fully loaded, resolved and validated configuration parameters.
type ConfigParameter struct {
	DataObjectWithReference
	Name        ParameterName          `yaml:"name"`
	Description generic.NonEmptyString `yaml:"description"`
	Value       ParameterValue         `yaml:"value"`
}
