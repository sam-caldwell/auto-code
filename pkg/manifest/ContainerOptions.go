package manifest

// ContainerOptions - optional parameters for a container definition
type ContainerOptions struct {
	DataObjectWithReference
	Publish     []uint16          `yaml:"publish,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
}
