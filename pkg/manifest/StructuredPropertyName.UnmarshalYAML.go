package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - unmarshal a YAML object into the structured property name
//
// This should accept a property string similar to the 'jq' program.
func (p *StructuredPropertyName) UnmarshalYAML(node *yaml.Node) error {
	const pattern = `^([a-zA-Z_][a-zA-Z0-9_-]*|\` +
		"`" + `[^\` + "`" + `]*` + "`" + `|\"[^\"]*\")(\[[0-9]+\])*$`

	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid property name (%s)", trimmedValue)
	}

	*p = StructuredPropertyName(value)
	return nil
}
