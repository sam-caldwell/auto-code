package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// ParseYamlObjectWithReferences - handle the unmarshalling of both $ref and non-$ref YAML dataContract.
//
// For example: To implement a $ref field, as in the below struct, we add DataObjectWithReference
//
//	  type ArtifactDescriptor struct {
//		   DataObjectWithReference
//		   Name         ArtifactName      `yaml:"name,omitempty"`
//		   Description  NonEmptyString    `yaml:"description,omitempty"`
//	  	Type         ArtifactType      `yaml:"type,omitempty"`
//		   Dependencies []ArtifactName    `yaml:"dependencies,omitempty"`
//		   Options      []ArtifactOptions `yaml:"options,omitempty"`
//	  }
//
// This, then, adds the methods needed to meet the DataObjectWithReferences interface.
// ToDo: detect loops
func ParseYamlObjectWithReferences(node *yaml.Node, dataObject DataObjectWithReferences) error {

	// decode the dataObject for its explicit dataContract
	if err := node.Decode(&dataObject); err != nil {
		return fmt.Errorf("failed to decode explicit dataContract: %v", err)
	}
	// if the dataObject has no external references, return the dataObject as-is
	if ref := dataObject.GetRef(); ref == nil {
		return nil

		// else, if we do have an external reference, load/parse/merge the content.
	} else {
		// Load the referenced YAML file
		fileContent, err := os.ReadFile(string(*ref))
		if err != nil {
			return fmt.Errorf("failed to read referenced file %s: %v", *ref, err)
		}

		// Unmarshal the referenced file into the ArtifactDescriptorDataObject
		if err := yaml.Unmarshal(fileContent, &dataObject); err != nil {
			return fmt.Errorf("failed to unmarshal referenced file %s: %v", ref, err)
		}
	}
	return nil
}
