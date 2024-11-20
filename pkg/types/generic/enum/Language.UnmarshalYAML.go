package enum

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal the YAML CPU architecture field.
func (l *Language) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch v := strings.ToLower(strings.TrimSpace(value)); v {
	case "golang":
		*l = Golang
	case "typescript-react":
		*l = TypeScriptReact
	case "sql":
		*l = Sql
	default:
		return errors.New("unknown architecture: " + v)
	}
	return nil
}
