package httpContract

import "github.com/sam-caldwell/auto-code/pkg/manifest/artifact"

// Root - defines an HTTP contract
type Root struct {
	Artifacts []artifact.Name          `yaml:"artifacts"`
	Endpoints []HttpEndpointDescriptor `yaml:"endpoints"`
}
