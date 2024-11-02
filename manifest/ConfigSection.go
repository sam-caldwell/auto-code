package manifest

import (
	"github.com/sam-caldwell/auto-code/data"
)

// ConfigSection - The Manifest.yaml Config section
type ConfigSection struct {
	Sources data.SourceList `yaml:"sources"`

	Properties ConfigProperties `yaml:"properties"`

	File ConfigFile `yaml:"file"`

	Environment ConfigEnvironment `yaml:"environment"`

	Commandline CommandlineArgList `yaml:"commandline"`
}
