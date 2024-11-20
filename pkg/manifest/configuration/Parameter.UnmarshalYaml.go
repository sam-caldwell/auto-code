package configuration

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Parse an Parameter and store its information properly.
//
// This method will parse the YAML Parameter and resolve any external $ref
// which may be included.
func (p *Parameter) UnmarshalYAML(node *yaml.Node) error {

	// Parse the ArtifactDescriptor, handling both $ref and non-$ref dataContract
	return manifest.ParseYamlObjectWithReferences(node, p)
}
