package manifest

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal the YAML object for ArtifactType as the enumerated type underlying it.
//
// $ref not supported
func (t *ArtifactType) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch v := strings.ToLower(strings.TrimSpace(value)); v {
	case "service":
		*t = service
	case "external":
		*t = external
	case "binary":
		*t = binary
	default:
		return errors.New("unknown architecture: " + v)
	}
	return nil
}
