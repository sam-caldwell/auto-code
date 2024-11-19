package manifest

// ConfigFile - represents a configuration file
type ConfigFile struct {
	DataObjectWithReference
	Name   FileName         `yaml:"name"`
	Format ConfigFileFormat `yaml:"format"`
	Schema ConfigFileSchema `yaml:"schema"`
}
