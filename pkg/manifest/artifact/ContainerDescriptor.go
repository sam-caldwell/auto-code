package artifact

import (
	"github.com/sam-caldwell/auto-code/pkg/types/cpu"
	"github.com/sam-caldwell/auto-code/pkg/types/docker"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// ContainerDescriptor - represents a docker container definition
type ContainerDescriptor struct {
	xref.DataObjectWithReference
	Name            docker.ContainerName    `yaml:"name"`
	CpuArchitecture cpu.Architecture        `yaml:"cpuArch"`
	Image           docker.Image            `yaml:"image"`
	Dockerfile      docker.FileName         `yaml:"dockerfile"`
	Options         docker.ContainerOptions `yaml:"options"`
}
