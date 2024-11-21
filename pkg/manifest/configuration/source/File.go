package source

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/file"
	"github.com/sam-caldwell/auto-code/pkg/types/generic/enum"
)

// File - represents a configuration file
type File struct {
	manifest.DataObjectWithReference
	Name   file.Name       `yaml:"name"`
	Format enum.FileFormat `yaml:"format"`
	Schema []FileSchema    `yaml:"schema"`
}
