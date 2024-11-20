package artifact

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/cpu"
	"github.com/sam-caldwell/auto-code/pkg/types/docker"
)

// ContainerDescriptor - represents a docker container definition
type ContainerDescriptor struct {
	manifest.DataObjectWithReference
	Name            docker.ContainerName    `yaml:"name"`
	CpuArchitecture cpu.CpuArchitecture     `yaml:"cpuArch"`
	Image           docker.Image            `yaml:"image"`
	Dockerfile      docker.FileName         `yaml:"dockerfile"`
	Options         docker.ContainerOptions `yaml:"options"`
}
