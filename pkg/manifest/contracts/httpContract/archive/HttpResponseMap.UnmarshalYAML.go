package archive

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strconv"
)

// UnmarshalYAML is a custom unmarshal function for HttpResponseMap
func (h *HttpResponseMap) UnmarshalYAML(node *yaml.Node) error {
	// Initialize the map
	*h = make(HttpResponseMap)

	// Iterate over the node content (which should be a sequence of key-value pairs)
	for i := 0; i < len(node.Content); i += 2 {
		// The first node is the key (status code), the second node is the value (HttpResponseDescriptor)
		statusCodeNode := node.Content[i]
		responseNode := node.Content[i+1]

		// Convert the status code from a string to uint16
		statusCode, err := strconv.ParseUint(statusCodeNode.Value, 10, 16)
		if err != nil {
			return fmt.Errorf("invalid status code %s: %v", statusCodeNode.Value, err)
		}

		// Create a HttpResponseDescriptor from the response node
		var response HttpResponseDescriptor
		if err := responseNode.Decode(&response); err != nil {
			return fmt.Errorf("failed to unmarshal HttpResponseDescriptor: %v", err)
		}

		// Add the key-value pair to the map
		(*h)[uint16(statusCode)] = response
	}

	return nil
}
