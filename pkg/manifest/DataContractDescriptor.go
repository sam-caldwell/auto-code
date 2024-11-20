package manifest

import "github.com/sam-caldwell/auto-code/pkg/types/generic"

// DataContractDescriptor - represents the content of a data contract
type DataContractDescriptor struct {
	DataObjectWithReference
	Name      generic.NonEmptyString `yaml:"name"`
	Artifacts []ArtifactName         `yaml:"artifacts"`
	Schema    DataSchemaDescriptors  `yaml:"schema"`
}
