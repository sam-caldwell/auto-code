package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - unmarshal and validates a US Copyright string from YAML
//
// Expected format: "(c) YEAR Name All Rights Reserved."
func (u *USCopyright) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^\(c\) \d{4} [a-zA-Z ,.'-]+ All Rights Reserved\.$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	// Trim extra whitespace before validation
	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid US Copyright format: %s", value)
	}

	*u = USCopyright(trimmedValue)
	return nil
}
