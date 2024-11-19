package manifest

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal a YAML object into ConfigFileFormat
func (cfg *ConfigFileFormat) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "yaml":
		*cfg = YAML
	case "json":
		*cfg = JSON
	default:
		return errors.New("unknown architecture: " + value)
	}
	return nil
}
