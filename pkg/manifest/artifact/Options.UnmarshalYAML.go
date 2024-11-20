package artifact

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Parse an Options and store its information properly.
//
// This method will parse the YAML Options and resolve any external $ref
// which may be included.
func (a *Options) UnmarshalYAML(node *yaml.Node) error {

	// Parse the ArtifactDescriptor, handling both $ref and non-$ref data
	return manifest.ParseYamlObjectWithReferences(node, a)

}
