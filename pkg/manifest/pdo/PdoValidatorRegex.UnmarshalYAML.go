package pdo

import (
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - unmarshal a yaml object which is expected to contain a regular expression.
func (p *PdoValidatorRegex) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	re, err := regexp.Compile(strings.TrimSpace(value))
	if err != nil {
		return err
	}
	if err = p.Set(re); err != nil {
		return err
	}
	return nil
}
