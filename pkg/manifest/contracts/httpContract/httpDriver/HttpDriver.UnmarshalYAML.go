package httpDriver

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - Unmarshal a yaml node to this HttpDriver object
func (e *HttpDriver) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "rest":
		*e = Rest
	case "graphql":
		*e = GraphQl
	default:
		return errors.New("unknown http driver: " + value)
	}
	return nil
}
