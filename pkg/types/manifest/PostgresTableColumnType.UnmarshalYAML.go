package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - Unmarshal a YAML object to PostgresTableColumnType enum
func (p *PostgresTableColumnType) UnmarshalYAML(node *yaml.Node) error {

	const pattern = `^[a-zA-Z][a-zA-Z0-9_]{0,62}(\([0-9]{1,20}\))?$`
	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid table name (%s)", trimmedValue)
	}

	*p = PostgresTableColumnType(value)
	return nil
}
