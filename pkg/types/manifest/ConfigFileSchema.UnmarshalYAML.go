package manifest

import "gopkg.in/yaml.v3"

// UnmarshalYAML - unmarshal a yaml object to a ConfigFileSchema object
func (cfg *ConfigFileSchema) UnmarshalYAML(node *yaml.Node) error {
	return ParseYamlObjectWithReferences(node, cfg)
}
