package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
)

// UnmarshalYAML - unmarshall and validate a container name
func (c *ContainerName) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^[a-zA-Z][a-zA-Z0-9\-_]{1,10}[a-zA-Z0-9]?$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	if re := regexp.MustCompile(pattern); !re.MatchString(value) {
		return fmt.Errorf("invalid container name: %s", c.String())
	}

	*c = ContainerName(value)

	return nil

}
