package source

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/manifest/configuration"
	"github.com/sam-caldwell/auto-code/pkg/types/file"
)

// File - represents a configuration file
type File struct {
	manifest.DataObjectWithReference
	Name   file.Name                `yaml:"name"`
	Format configuration.FileFormat `yaml:"format"`
	Schema configuration.FileSchema `yaml:"schema"`
}
