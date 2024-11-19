package manifest

// ConfigFileSchema - the data schema for a configuration file
type ConfigFileSchema struct {
	DataObjectWithReference
	Path      StructuredPropertyName `yaml:"path"`
	Parameter ParameterName          `yaml:"parameter"`
}
