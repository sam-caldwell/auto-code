package manifest

import "github.com/sam-caldwell/auto-code/pkg/types/generic"

// HttpResponseDescriptor - represents an HTTP response
type HttpResponseDescriptor struct {
	DataObjectWithReference
	Description generic.NonEmptyString `yaml:"description"`
	ContentType HttpContentType        `yaml:"content_type"`
	Headers     HttpResponseHeader     `yaml:"headers"`
	Body        HttpResponseBody       `yaml:"body"`
}
