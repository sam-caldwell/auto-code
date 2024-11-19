package manifest

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - Unmarshal a yaml object to HttpEndpointMethod
func (h *HttpEndpointMethod) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "get":
		*h = GET
	case "post":
		*h = POST
	case "put":
		*h = PUT
	case "delete":
		*h = DELETE
	default:
		return errors.New("unknown architecture: " + value)
	}
	return nil
}
