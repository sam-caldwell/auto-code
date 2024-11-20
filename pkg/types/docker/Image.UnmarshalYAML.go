package docker

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
)

// UnmarshalYAML - unmarshal and validates a container image from YAML
func (ci *ContainerImage) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^[a-zA-Z0-9\-_]+(?:/[a-zA-Z0-9\-_]+)*(?::[a-zA-Z0-9\-.]+)?$`
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}

	if re := regexp.MustCompile(pattern); !re.MatchString(value) {
		return fmt.Errorf("invalid container image name: %s", value)
	}

	*ci = ContainerImage(value)
	return nil
}
