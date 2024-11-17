package manifest

type ArtifactOptions struct {
	Language  Language            `yaml:"language,omitempty"`
	Sources   PosixDirectory      `yaml:"sources,omitempty"`
	Url       UrlPattern          `yaml:"url,omitempty"`
	Container ContainerDescriptor `yaml:"container,omitempty"`
}
