package manifest

// DataContractDescriptor - represents the content of a data contract
type DataContractDescriptor struct {
	DataObjectWithReference
	Name      NonEmptyString        `yaml:"name"`
	Artifacts []ArtifactName        `yaml:"artifacts"`
	Schema    DataSchemaDescriptors `yaml:"schema"`
}
