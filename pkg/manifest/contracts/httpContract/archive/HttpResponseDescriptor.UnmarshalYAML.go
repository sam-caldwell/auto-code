package httpContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Unmarshal a yaml node into a HttpResponseDescriptor
func (h *HttpResponseDescriptor) UnmarshalYAML(node *yaml.Node) error {
	return manifest.ParseYamlObjectWithReferences(node, h)
}
