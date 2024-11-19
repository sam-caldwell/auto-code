package manifest

import (
	"errors"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - unmarshal the YAML CPU architecture field.
func (cpu *CpuArchitecture) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch value {
	case "arm":
		*cpu = arm
	case "arm64":
		*cpu = arm64
	case "amd64":
		*cpu = amd64
	default:
		return errors.New("unknown architecture: " + value)
	}
	return nil
}
