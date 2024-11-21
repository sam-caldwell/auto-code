package source

import (
	"github.com/sam-caldwell/auto-code/pkg/types/file"
	"github.com/sam-caldwell/auto-code/pkg/types/generic/enum"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// File - represents a configuration file
type File struct {
	xref.DataObjectWithReference
	Name   file.Name       `yaml:"name"`
	Format enum.FileFormat `yaml:"format"`
	Schema []FileSchema    `yaml:"schema"`
}
