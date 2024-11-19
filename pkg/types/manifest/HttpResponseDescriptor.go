package manifest

// HttpResponseDescriptor - represents an HTTP response
type HttpResponseDescriptor struct {
	DataObjectWithReference
	Description NonEmptyString     `yaml:"description"`
	ContentType HttpContentType    `yaml:"content_type"`
	Headers     HttpResponseHeader `yaml:"headers"`
	Body        HttpResponseBody   `yaml:"body"`
}
