package docker

import "github.com/sam-caldwell/auto-code/pkg/manifest"

// ContainerOptions - optional parameters for a container definition
type ContainerOptions struct {
	manifest.DataObjectWithReference
	Publish     []uint16          `yaml:"publish,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
}
