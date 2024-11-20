package manifest

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/artifact"
	"github.com/sam-caldwell/auto-code/pkg/manifest/configuration"
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/dataContract"
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/httpContract"
	"github.com/sam-caldwell/auto-code/pkg/types/version"
)

// Root - represents the top-level of the auto-code YAML specification
//
// Note: The document root does not support $ref because this would be
//
//	confusing.
type Root struct {
	Version       version.Version             `yaml:"auto-code"`
	Info          Info                        `yaml:"info"`
	Artifacts     []artifact.Descriptor       `yaml:"artifacts"`
	Configuration configuration.Configuration `yaml:"configuration"`
	HttpContract  httpContract.Root           `yaml:"httpContract"`
	DataContract  dataContract.Root           `yaml:"dataContract"`
}
