// Package manifest generator/Manifest.go
package manifest

import (
	data2 "github.com/sam-caldwell/auto-code/data"
)

// Commandline - maps CLI arguments to properties
type Commandline struct {
	Required bool `yaml:"required"`

	Short data2.ShortArgumentString `yaml:"short,omitempty"`

	Long data2.LongArgumentString `yaml:"long,omitempty"`
}
