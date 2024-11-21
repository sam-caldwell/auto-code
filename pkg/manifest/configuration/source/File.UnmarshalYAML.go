package source

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - unmarshal the YAML object into the File object
func (cfg *File) UnmarshalYAML(node *yaml.Node) error {
	return manifest.ParseYamlObjectWithReferences(node, cfg)
}
