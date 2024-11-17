package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - Unmarshal a string and validate its value
func (n *ArtifactName) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^[a-zA-Z][a-zA-Z0-9\-_\.]{0,14}[a-zA-Z0-9]$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid artifact name (%s)", trimmedValue)
	}

	*n = ArtifactName(value)
	return nil
}
