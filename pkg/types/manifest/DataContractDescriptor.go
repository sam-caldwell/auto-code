package manifest

// DataContractDescriptor - represents the content of a data contract
type DataContractDescriptor struct {
	DataObjectWithReference
	Name      NonEmptyString       `yaml:"name"`
	Artifacts []ArtifactName       `yaml:"artifacts"` // Artifacts - Map the data contract to one or more artifacts
	Schema    DataSchemaDescriptor `yaml:"schema"`
}
