package httpContract

import "github.com/sam-caldwell/auto-code/pkg/manifest/artifact"

type Root struct {
	Artifacts []artifact.Name          `yaml:"artifacts"`
	Endpoints []HttpEndpointDescriptor `yaml:"endpoints"`
}
