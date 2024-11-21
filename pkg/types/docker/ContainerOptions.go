package docker

import (
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// ContainerOptions - optional parameters for a container definition
type ContainerOptions struct {
	xref.DataObjectWithReference
	Publish     []uint16          `yaml:"publish,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
}
