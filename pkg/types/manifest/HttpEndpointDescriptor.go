package manifest

// HttpEndpointDescriptor - Represents an HTTP endpoint
type HttpEndpointDescriptor struct {
	DataObjectWithReference
	Path        HttpEndpointName    `yaml:"path"`
	Version     HttpEndpointVersion `yaml:"version"`
	Method      HttpEndpointMethod  `yaml:"method"`
	Description NonEmptyString      `yaml:"description"`
	Decorators  []DecoratorName     `yaml:"decorators"`
	Function    HttpHandlerFunction `yaml:"function"`
	Responses   HttpResponseMap     `yaml:"responses"`
}
