package manifest

import "gopkg.in/yaml.v3"

// UnmarshalYAML - unmarshal the YAML object into the CommandLineArg object
func (cfg *CommandLineArg) UnmarshalYAML(node *yaml.Node) error {
	return ParseYamlObjectWithReferences(node, cfg)
}
