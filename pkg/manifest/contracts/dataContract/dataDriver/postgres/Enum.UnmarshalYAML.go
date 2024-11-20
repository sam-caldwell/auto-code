package postgres

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
)

// UnmarshalYAML - unmarshal yaml object to an enumerated type
func (e *Enum) UnmarshalYAML(node *yaml.Node) error {

	const identifierPattern = `^[a-zA-Z][a-zA-Z0-9_]{0,14}[a-zA-Z0-9]$`

	var value Enum
	if err := node.Decode(&value); err != nil {
		return err
	}

	if re := regexp.MustCompile(identifierPattern); !re.MatchString(value.Name.String()) {
		return fmt.Errorf("invalid enum name identifier: %s", value.Name)
	}

	for _, element := range value.Elements {
		if re := regexp.MustCompile(identifierPattern); !re.MatchString(element.String()) {
			return fmt.Errorf("invalid enum element: %s", element)
		}
	}

	*e = Enum(value)
	return nil
}
