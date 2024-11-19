package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - Unmarshal YAML objects into a PostgresTableIndexName
func (p *PostgresTableIndexName) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^[a-zA-Z][a-zA-Z0-9_]{0,14}[a-zA-Z0-9]$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid table name (%s)", trimmedValue)
	}

	*p = PostgresTableIndexName(value)
	return nil
}
