package archive

import (
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// HttpResponseDescriptor - represents an HTTP response
type HttpResponseDescriptor struct {
	xref.DataObjectWithReference
	Description generic.NonEmptyString `yaml:"description"`
	ContentType ContentType            `yaml:"content_type"`
	Headers     HttpResponseHeader     `yaml:"headers"`
	Body        HttpResponseBody       `yaml:"body"`
}
