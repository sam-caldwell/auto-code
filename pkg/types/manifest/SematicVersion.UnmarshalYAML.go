package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - unmarshal and validates a semantic version string from YAML
func (sv *SemanticVersion) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^v(\d+)\.(\d+)\.(\d+)(?:-[a-zA-Z0-9.-]+)?(?:\+[a-zA-Z0-9.-]+)?$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	// Trim whitespace from the input
	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid semantic version: %s", value)
	}

	*sv = SemanticVersion(trimmedValue)
	return nil
}
