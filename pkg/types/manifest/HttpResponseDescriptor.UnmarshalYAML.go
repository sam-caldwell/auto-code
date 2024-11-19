package manifest

import "gopkg.in/yaml.v3"

// UnmarshalYAML - Unmarshal a yaml node into a HttpResponseDescriptor
func (h *HttpResponseDescriptor) UnmarshalYAML(node *yaml.Node) error {
	return ParseYamlObjectWithReferences(node, h)
}
