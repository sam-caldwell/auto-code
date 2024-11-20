package license

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal the YAML CPU architecture field.
func (l *Name) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch v := strings.ToLower(strings.TrimSpace(value)); v {
	case "mit":
		*l = MIT
	case "apache2":
		*l = Apache2
	case "bsd":
		*l = BSD
	case "", "proprietary":
		*l = Proprietary
	default:
		return fmt.Errorf("unknown/unrecognized Name '%s'", v)
	}
	return nil
}
