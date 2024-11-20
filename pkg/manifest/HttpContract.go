package manifest

import "github.com/sam-caldwell/auto-code/pkg/manifest/artifact"

type HttpContract struct {
	Artifacts []artifact.Name          `yaml:"artifacts"`
	Endpoints []HttpEndpointDescriptor `yaml:"endpoints"`
}
