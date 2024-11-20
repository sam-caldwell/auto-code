package manifest

import "github.com/sam-caldwell/auto-code/pkg/types/directory"

// ArtifactOptions - represents the options associated with a project artifact
type ArtifactOptions struct {
	DataObjectWithReference
	Language  Language            `yaml:"language,omitempty"`
	Sources   directory.Name      `yaml:"sources,omitempty"`
	Url       UrlPattern          `yaml:"url,omitempty"`
	Container ContainerDescriptor `yaml:"container,omitempty"`
}
