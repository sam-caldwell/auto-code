package manifest

import "gopkg.in/yaml.v3"

// UnmarshalYAML - unmarshal the YAML object into the ConfigFile object
func (cfg *ConfigFile) UnmarshalYAML(node *yaml.Node) error {
	return ParseYamlObjectWithReferences(node, cfg)
}
