package configuration

import (
	"gopkg.in/yaml.v3"
)

func (c *Descriptor) UnmarshalYAML(node *yaml.Node) error {

	if err := node.Decode(c); err != nil {
		return err
	}

	if err := c.Validate(); err != nil {
		return err
	}

	return nil
}
