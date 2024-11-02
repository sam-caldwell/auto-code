// Package manifest generator/Manifest.go
package manifest

import (
	"github.com/sam-caldwell/auto-code/data"
	"github.com/sam-caldwell/types"
)

// ConfigFile defines file-based configuration
type ConfigFile struct {
	File types.FileNameString `yaml:"file"`

	Required bool `yaml:"required,omitempty"`

	Format data.FileFormatString `yaml:"format"`

	Map PropertyMap `yaml:"map"`
}
