package email

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
)

// UnmarshalYAML - unmarshal and validates an email address from YAML (RFC5322)
func (e *Address) UnmarshalYAML(node *yaml.Node) error {
	// Regex pattern for a valid email address (simplified version)
	const pattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	if re := regexp.MustCompile(pattern); !re.MatchString(value) {
		return fmt.Errorf("invalid email address: %s", value)
	}

	*e = Address(value)
	return nil
}
