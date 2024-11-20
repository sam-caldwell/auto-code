package artifact

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/directory"
	"github.com/sam-caldwell/auto-code/pkg/types/generic/enum"
	"github.com/sam-caldwell/auto-code/pkg/types/url"
)

// Options - represents the options associated with a project artifact
type Options struct {
	manifest.DataObjectWithReference
	Language  enum.Language                `yaml:"language,omitempty"`
	Sources   directory.Name               `yaml:"sources,omitempty"`
	Url       url.Pattern                  `yaml:"url,omitempty"`
	Container manifest.ContainerDescriptor `yaml:"container,omitempty"`
}
