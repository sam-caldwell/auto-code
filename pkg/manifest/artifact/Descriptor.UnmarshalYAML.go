package artifact

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Parse a Descriptor and store its information properly.
//
// This method will parse the YAML Descriptor and resolve any external $ref
// which may be included.
func (a *Descriptor) UnmarshalYAML(node *yaml.Node) error {

	// Parse the Descriptor, handling both $ref and non-$ref data
	return manifest.ParseYamlObjectWithReferences(node, a)

}
