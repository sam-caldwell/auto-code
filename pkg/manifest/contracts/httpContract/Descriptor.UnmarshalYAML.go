package httpContract

import "gopkg.in/yaml.v3"

// UnmarshalYAML - Unmarshal yaml object into the descriptor object
func (d *Descriptor) UnmarshalYAML(node *yaml.Node) error {
	if err := node.Decode(&d); err != nil {
		return err
	}
	return nil
}
