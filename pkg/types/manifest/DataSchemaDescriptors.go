package manifest

import "gopkg.in/yaml.v3"

// DataSchemaDescriptors - an interface for a data contract schema
type DataSchemaDescriptors interface {
	String() string
	UnmarshalYAML(node *yaml.Node) error
}
