package configuration

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/configuration/source"
	"github.com/sam-caldwell/auto-code/pkg/types/generic/enum"
)

// Configuration - define the standardized mechanism by which the solution is configured.
//
// This struct creates a standardized configuration mechanism for the solution (all artifacts)
// in order to ensure consistency across artifacts and a positive user experience.  For example,
// this feature allows all artifacts to be defined by configuration file information, then to
// overlay environment variables and command-line arguments so that the applications themselves
// can reference the resulting dataContract using a common parameter name scheme without a lot of effort.
type Configuration struct {
	MergeOrder  []enum.MergeOrder       `yaml:"merge-order,omitempty"`
	Parameters  []Parameter             `yaml:"parameters,omitempty"`
	Files       []source.File           `yaml:"files,omitempty"`
	Environment []source.Environment    `yaml:"environment,omitempty"`
	CommandLine []source.CommandLineArg `yaml:"command-line,omitempty"`
}
