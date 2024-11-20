package manifest

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/artifact"
	"github.com/sam-caldwell/auto-code/pkg/types/version"
)

// Root - represents the top-level of the auto-code YAML specification
//
// Note: The document root does not support $ref because this would be
//
//	confusing.
type Root struct {
	Version       version.Version       `yaml:"auto-code"`
	Info          Info                  `yaml:"info"`
	Artifacts     []artifact.Descriptor `yaml:"artifacts"`
	Configuration Configuration         `yaml:"configuration"`
	HttpContract  HttpContract          `yaml:"httpContract"`
	DataContract  DataContract          `yaml:"dataContract"`
}
