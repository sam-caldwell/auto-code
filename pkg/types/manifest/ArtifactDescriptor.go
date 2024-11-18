package manifest

// ArtifactDescriptor - represents a single project artifact.
//
// An artifact may be explicitly defined or may reference an external document
// which will be consumed to populate the data.
type ArtifactDescriptor struct {
	DataObjectWithReference
	Name         ArtifactName      `yaml:"name,omitempty"`
	Description  NonEmptyString    `yaml:"description,omitempty"`
	Type         ArtifactType      `yaml:"type,omitempty"`
	Dependencies []ArtifactName    `yaml:"dependencies,omitempty"`
	Options      []ArtifactOptions `yaml:"options,omitempty"`
}
