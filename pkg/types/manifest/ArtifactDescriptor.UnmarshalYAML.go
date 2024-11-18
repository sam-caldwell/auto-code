package manifest

import (
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Parse an ArtifactDescriptor and store its information properly.
//
// This method will parse the YAML ArtifactDescriptor and resolve any external $ref
// which may be included.
func (a *ArtifactDescriptor) UnmarshalYAML(node *yaml.Node) error {

	// Parse the ArtifactDescriptor, handling both $ref and non-$ref data
	return ParseYamlObjectWithReferences(node, a)
}
