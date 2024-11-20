package manifest

// ContainerDescriptor - represents a docker container definition
type ContainerDescriptor struct {
	DataObjectWithReference
	Name            ContainerName    `yaml:"name"`
	CpuArchitecture CpuArchitecture  `yaml:"cpuArch"`
	Image           ContainerImage   `yaml:"image"`
	Dockerfile      DockerFileName   `yaml:"dockerfile"`
	Options         ContainerOptions `yaml:"options"`
}
