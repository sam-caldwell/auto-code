package generic

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal and validates a non-empty string from YAML
//
// $ref not supported
func (n *NonEmptyString) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.TrimSpace(value)
	if len(trimmedValue) == 0 {
		return fmt.Errorf("string value cannot be empty")
	}

	*n = NonEmptyString(trimmedValue)
	return nil
}
