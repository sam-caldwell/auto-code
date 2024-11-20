package dataCommon

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - Unmarshal a YAML node into a PostgresTableName object
func (p *Identifier) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^[a-zA-Z][a-zA-Z0-9_]{0,14}[a-zA-Z0-9]$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid table name (%s)", trimmedValue)
	}

	*p = Identifier(value)
	return nil
}
