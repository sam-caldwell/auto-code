package manifest

import "gopkg.in/yaml.v3"

// UnmarshalYAML - unmarshal the YAML object into the EnvironmentVariable object
func (cfg *EnvironmentVariable) UnmarshalYAML(node *yaml.Node) error {
	return ParseYamlObjectWithReferences(node, cfg)
}
