package file

import (
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - unmarshal a YAML object into the Name string
func (f *Name) UnmarshalYAML(node *yaml.Node) error {

	if err := node.Decode(f); err != nil {
		return err
	}

	if err := f.Validate(); err != nil {
		return err
	}

	return nil

}
