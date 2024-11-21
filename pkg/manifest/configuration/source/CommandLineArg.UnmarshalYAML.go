package source

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - unmarshal the YAML object into the CommandLineArg object
func (cfg *CommandLineArg) UnmarshalYAML(node *yaml.Node) error {
	return manifest.ParseYamlObjectWithReferences(node, cfg)
}
