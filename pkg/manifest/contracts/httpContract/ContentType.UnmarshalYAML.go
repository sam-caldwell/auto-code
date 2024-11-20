package httpContract

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
)

// UnmarshalYAML - Unmarshal a yaml node to an ContentType object
func (h *ContentType) UnmarshalYAML(node *yaml.Node) error {

	const pattern = `` +
		`^[a-zA-Z0-9!#$&'*+.^_-]+` +
		`/[a-zA-Z0-9!#$&'*+.^_-]+` +
		`(;[a-zA-Z0-9!#$&'*+.^_-]+=([a-zA-Z0-9!#$&'*+.^_-]+|"[^"]*"))*$`

	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	// Check if the given value matches the required format
	if re := regexp.MustCompile(pattern); !re.MatchString(value) {
		return fmt.Errorf("invalid httpContract Content-Type header string: %s", value)
	}

	*h = ContentType(value)

	return nil
}
