package manifest

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/artifact"
	"github.com/sam-caldwell/auto-code/pkg/manifest/configuration"
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts"
	"github.com/sam-caldwell/auto-code/pkg/manifest/metadata"
	"github.com/sam-caldwell/auto-code/pkg/types/version"
)

// Root - represents the top-level of the auto-code YAML specification
//
// Note: The document root does not support $ref because this would be
//
//	confusing.
//
// ToDo: verify that the dataContract, httpContract names are in the artifacts list.
type Root struct {
	Version       version.Version          `yaml:"auto-code"`
	Info          metadata.Info            `yaml:"info"`
	Artifacts     []artifact.Descriptor    `yaml:"artifacts"`
	Configuration configuration.Descriptor `yaml:"configuration"`
	Contracts     contracts.Descriptor     `yaml:"contracts"`
}
