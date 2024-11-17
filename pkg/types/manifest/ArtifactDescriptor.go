package manifest

type ArtifactDescriptor struct {
	Name         ArtifactName      `yaml:"name"`
	Description  NonEmptyString    `yaml:"description"`
	Type         ArtifactType      `yaml:"type"`
	Dependencies []ArtifactName    `yaml:"dependencies,omitempty"`
	Options      []ArtifactOptions `yaml:"options,omitempty"`
}
