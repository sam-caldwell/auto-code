package configuration

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - unmarshal the YAML object into the EnvironmentVariable object
func (cfg *EnvironmentVariable) UnmarshalYAML(node *yaml.Node) error {
	return manifest.ParseYamlObjectWithReferences(node, cfg)
}
