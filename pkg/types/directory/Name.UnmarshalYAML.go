package directory

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - unmarshal and validates a POSIX directory path from YAML
// ToDo: add windows support later
func (p *Name) UnmarshalYAML(node *yaml.Node) error {
	// Regular expression to validate POSIX directory paths
	const pattern = `^(/?([a-zA-Z0-9_\-\.]+/)*[a-zA-Z0-9_\-\.]+/?)$`

	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid POSIX directory path: %s", trimmedValue)
	}

	*p = Name(trimmedValue)
	return nil
}
