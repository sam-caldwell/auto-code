package manifest

import (
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Parse an ArtifactDescriptor and store its information properly.
//
// This method will parse the YAML ArtifactDescriptor and resolve any external $ref
// which may be included.
func (d *ArtifactDescriptor) UnmarshalYAML(node *yaml.Node) error {

	var dataObject ArtifactDescriptorDataObject
	// Parse the ArtifactDescriptor, handling both $ref and non-$ref data
	if data, err := ParseYamlObject(node, dataObject, &d.Ref); err != nil {
		return err
	} else {
		// Store the parsed data into the ArtifactDescriptor
		d.Data = data.(ArtifactDescriptorDataObject)
		return nil
	}
}
