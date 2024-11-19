package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
)

// UnmarshalYAML - Unmarshal a YAML object into an HttpEndpointPath
func (h *HttpEndpointPath) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^\/([a-zA-Z0-9-._~!$&'()*+,;=:@%]+(\/[a-zA-Z0-9-._~!$&'()*+,;=:@%]*)*)?$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	// Check if the given value matches the required format
	if re := regexp.MustCompile(pattern); !re.MatchString(value) {
		return fmt.Errorf("invalid http path: %s", value)
	}

	*h = HttpEndpointPath(value)

	return nil
}
