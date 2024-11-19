package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - Unmarshal a yaml object to the HttpDecoratorName
func (h *HttpDecoratorName) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^[a-zA-Z][a-zA-Z0-9][a-zA-Z0-9]$`

	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid HTTP decorator name (%s)", trimmedValue)
	}

	*h = HttpDecoratorName(trimmedValue)

	return nil
}
