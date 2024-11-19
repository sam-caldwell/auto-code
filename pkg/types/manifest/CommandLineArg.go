package manifest

// CommandLineArg - represents a command-line argument for the Configuration object
type CommandLineArg struct {
	DataObjectWithReference
	Path      StructuredPropertyName `yaml:"path"`
	Parameter ParameterName          `yaml:"parameter"`
}
