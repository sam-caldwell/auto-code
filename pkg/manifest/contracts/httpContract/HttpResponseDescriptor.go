package httpContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// HttpResponseDescriptor - represents an HTTP response
type HttpResponseDescriptor struct {
	manifest.DataObjectWithReference
	Description generic.NonEmptyString `yaml:"description"`
	ContentType ContentType            `yaml:"content_type"`
	Headers     HttpResponseHeader     `yaml:"headers"`
	Body        HttpResponseBody       `yaml:"body"`
}
