package dataContract

import "gopkg.in/yaml.v3"

// SchemaDescriptors - an interface for a dataContract contract schema
type SchemaDescriptors interface {
	String() string
	UnmarshalYAML(node *yaml.Node) error
}
