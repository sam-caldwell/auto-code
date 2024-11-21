package artifact

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal the YAML object for Type as the enumerated type underlying it.
//
// $ref not supported
func (t *Type) UnmarshalYAML(node *yaml.Node) error {

	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	switch v := strings.ToLower(strings.TrimSpace(value)); v {
	case "service":
		*t = Service
	case "external":
		*t = External
	case "binary":
		*t = Binary
	default:
		return errors.New("unknown architecture: " + v)
	}

	return nil
}
