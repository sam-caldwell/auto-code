package contracts

import "gopkg.in/yaml.v3"

// SchemaDescriptors - an interface for a contract schema (applies to both http and data contracts)
type SchemaDescriptors interface {

	// String - return a string representation of the object
	String() string

	// UnmarshalYAML - unmarshal the current YAML node to the given object
	UnmarshalYAML(node *yaml.Node) error

	// Generate - Generate object code based on the current state of the object
	Generate() error
}
