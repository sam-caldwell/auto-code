package manifest

// Configuration - define the standardized mechanism by which the solution is configured.
//
// This struct creates a standardized configuration mechanism for the solution (all artifacts)
// in order to ensure consistency across artifacts and a positive user experience.  For example,
// this feature allows all artifacts to be defined by configuration file information, then to
// overlay environment variables and command-line arguments so that the applications themselves
// can reference the resulting data using a common parameter name scheme without a lot of effort.
type Configuration struct {
	MergeOrder  []MergeOrder          `yaml:"merge-order,omitempty"`
	Parameters  []ConfigParameter     `yaml:"parameters,omitempty"`
	Files       []ConfigFile          `yaml:"files,omitempty"`
	Environment []EnvironmentVariable `yaml:"environment,omitempty"`
	CommandLine []CommandLineArg      `yaml:"command-line,omitempty"`
}
