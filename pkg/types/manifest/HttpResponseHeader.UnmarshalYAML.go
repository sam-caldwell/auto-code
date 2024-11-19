package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Unmarshal YAML node into an HttpResponseBody object
func (h *HttpResponseHeader) UnmarshalYAML(node *yaml.Node) error {
	// Ensure the map is initialized
	*h = make(HttpResponseHeader)

	// Check for an even number of elements (key-value pairs)
	if len(node.Content)%2 != 0 {
		return fmt.Errorf("invalid YAML format: expected an even number of elements (key-value pairs)")
	}

	// Process the key-value pairs
	for i := 0; i < len(node.Content); i += 2 {
		keyNode := node.Content[i]     // Key is at even index
		valueNode := node.Content[i+1] // Value is at odd index

		// Attempt to decode the value into a generic 'any' type
		var response string
		if err := valueNode.Decode(&response); err != nil {
			return fmt.Errorf("failed to unmarshal response value for key %s: %v", keyNode.Value, err)
		}

		// Add the decoded value to the map with the corresponding key
		(*h)[keyNode.Value] = response
	}

	return nil
}
