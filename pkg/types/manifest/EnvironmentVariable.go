package manifest

// EnvironmentVariable - represents an environment variable for the Configuration object
type EnvironmentVariable struct {
	DataObjectWithReference
	Path      StructuredPropertyName `yaml:"path"`
	Parameter ParameterName          `yaml:"parameter"`
}
