package manifest

// ArtifactDescriptor - represents a single project artifact.
//
// An artifact may be explicitly defined or may reference an external document
// which will be consumed to populate the data.
type ArtifactDescriptor struct {
	Ref  DocumentReference `yaml:"$ref"`
	Data ArtifactDescriptorDataObject
}

type ArtifactDescriptorDataObject struct {
	Name         ArtifactName      `yaml:"name"`
	Description  NonEmptyString    `yaml:"description"`
	Type         ArtifactType      `yaml:"type"`
	Dependencies []ArtifactName    `yaml:"dependencies,omitempty"`
	Options      []ArtifactOptions `yaml:"options,omitempty"`
}
