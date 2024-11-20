package manifest

import "github.com/sam-caldwell/auto-code/pkg/types/file"

// ConfigFile - represents a configuration file
type ConfigFile struct {
	DataObjectWithReference
	Name   file.Name        `yaml:"name"`
	Format ConfigFileFormat `yaml:"format"`
	Schema ConfigFileSchema `yaml:"schema"`
}
