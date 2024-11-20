package artifact

import (
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Unmarshal a string and validate its value
//
// $ref not supported
func (n *Name) UnmarshalYAML(node *yaml.Node) error {

	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	if err := n.Validate(); err != nil {
		return err
	}

	*n = Name(value)
	return nil
}
