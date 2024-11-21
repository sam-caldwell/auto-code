package source

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - unmarshal a yaml object to a FileSchema object
func (cfg *FileSchema) UnmarshalYAML(node *yaml.Node) error {
	return manifest.ParseYamlObjectWithReferences(node, cfg)
}
