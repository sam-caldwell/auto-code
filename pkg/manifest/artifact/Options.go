package artifact

import (
	"github.com/sam-caldwell/auto-code/pkg/types/directory"
	"github.com/sam-caldwell/auto-code/pkg/types/generic/enum"
	"github.com/sam-caldwell/auto-code/pkg/types/url"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// Options - represents the options associated with a project artifact
type Options struct {
	xref.DataObjectWithReference
	Language  enum.Language       `yaml:"language,omitempty"`
	Sources   directory.Name      `yaml:"sources,omitempty"`
	Url       url.Pattern         `yaml:"url,omitempty"`
	Container ContainerDescriptor `yaml:"container,omitempty"`
}
