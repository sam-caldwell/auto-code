package pdo

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - unmarshal a given YAML node into the associated ParameterDataObject.
func (pdo *ParameterDataObject) UnmarshalYAML(node *yaml.Node) error {

	return manifest.ParseYamlObjectWithReferences(node, pdo)

}
