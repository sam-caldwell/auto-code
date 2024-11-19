package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
)

// UnmarshalYAML - Unmarshal a yaml node to an HttpContentType object
func (h *HttpContentType) UnmarshalYAML(node *yaml.Node) error {

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
		return fmt.Errorf("invalid http Content-Type header string: %s", value)
	}

	*h = HttpContentType(value)

	return nil
}
