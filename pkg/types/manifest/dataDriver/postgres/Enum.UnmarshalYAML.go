package postgres

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
)

// UnmarshalYAML - unmarshal yaml object to an enumerated type
func (e *Enum) UnmarshalYAML(node *yaml.Node) error {
	// pattern - element verification
	const pattern = `[a-zA-Z][a-zA-Z0-9_][a-zA-Z0-9]`
	var value []string
	if err := node.Decode(&value); err != nil {
		return err
	}

	for _, element := range value {
		if re := regexp.MustCompile(pattern); !re.MatchString(element) {
			return fmt.Errorf("invalid enum element: %s", element)
		}
	}

	*e = Enum(value)
	return nil
}
