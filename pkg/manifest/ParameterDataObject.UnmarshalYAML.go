package manifest

import "gopkg.in/yaml.v3"

// UnmarshalYAML - unmarshal a given YAML node into the associated ParameterDataObject.
func (pdo *ParameterDataObject) UnmarshalYAML(node *yaml.Node) error {

	return ParseYamlObjectWithReferences(node, pdo)

}
