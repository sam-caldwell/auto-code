package generic

import (
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Unmarshal a YAML node into a PostgresTableName object
func (p *Identifier) UnmarshalYAML(node *yaml.Node) error {

	if err := node.Decode(p); err != nil {
		return err
	}

	if err := p.Validate(); err != nil {
		return err
	}

	return nil
}
