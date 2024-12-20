package cpu

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal the YAML CPU architecture field.
func (cpu *Architecture) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch strings.ToLower(strings.TrimSpace(value)) {
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
