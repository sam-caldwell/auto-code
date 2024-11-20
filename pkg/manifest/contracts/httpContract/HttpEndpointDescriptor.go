package httpContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// HttpEndpointDescriptor - Represents an HTTP endpoint
type HttpEndpointDescriptor struct {
	manifest.DataObjectWithReference
	Path        HttpEndpointPath       `yaml:"path"`
	Version     HttpEndpointVersion    `yaml:"version"`
	Method      HttpEndpointMethod     `yaml:"method"`
	Description generic.NonEmptyString `yaml:"description"`
	Decorators  []HttpDecoratorName    `yaml:"decorators"`
	Function    HttpHandlerFunction    `yaml:"function"`
	Responses   HttpResponseMap        `yaml:"responses"`
}