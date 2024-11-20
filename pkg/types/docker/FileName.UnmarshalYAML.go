package docker

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
)

// UnmarshalYAML - unmarshal and validate a Dockerfile name from YAML
//
// Valid patterns:
// 1. Dockerfile
// 2. <filename>.docker where <filename> follows POSIX path rules
func (d *FileName) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^(Dockerfile|([a-zA-Z0-9_\-./]+\.docker))$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	// Check if the given value matches the required format
	if re := regexp.MustCompile(pattern); !re.MatchString(value) {
		return fmt.Errorf("invalid Dockerfile name: %s", value)
	}

	*d = FileName(value)
	return nil
}
