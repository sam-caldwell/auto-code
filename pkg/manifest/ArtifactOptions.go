package manifest

// ArtifactOptions - represents the options associated with a project artifact
type ArtifactOptions struct {
	DataObjectWithReference
	Language  Language            `yaml:"language,omitempty"`
	Sources   PosixDirectory      `yaml:"sources,omitempty"`
	Url       UrlPattern          `yaml:"url,omitempty"`
	Container ContainerDescriptor `yaml:"container,omitempty"`
}
