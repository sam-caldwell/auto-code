package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal a given yaml node into MergeOrder
// It supports case-insensitive values and trims spaces before comparison.
// Recognized values are: "files", "environment", "command-line".
// Returns an error if the value is unrecognized.
func (m *MergeOrder) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch v := strings.ToLower(strings.TrimSpace(value)); v {
	case "files":
		*m = Files
	case "environment":
		*m = EnvironmentVars
	case "command-line":
		*m = CommandLine
	default:
		return fmt.Errorf("unknown/unexpected MergeOrder value (%s)", v)
	}
	return nil
}
