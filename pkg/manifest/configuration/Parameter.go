package configuration

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// Parameter - represents the top-level object in Descriptor.
//
// Given one or more data sources to configure a solution, the Descriptor object
// reads the configuration sources then merges them (using MergeOrder) before storing
// and validating the final value against the Descriptor.Parameter array.
//
// This object represents one element of the Descriptor.Parameter array, and it
// maps a parameter's name to its description and value (state and validator logic).
//
// When completed, the configuration process will expose the Parameter object as the
// single source of truth.
type Parameter struct {
	manifest.DataObjectWithReference
	Name        generic.Identifier     `yaml:"name"`
	Description generic.NonEmptyString `yaml:"description"`
	Value       ParameterValue         `yaml:"value"`
}
