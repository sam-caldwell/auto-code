package manifest

import (
	"gopkg.in/yaml.v3"
)

// String - return a string representation of HttpResponseHeader
func (h *HttpResponseHeader) String() string {
	result, err := yaml.Marshal(h)
	if err != nil {
		panic(err)
	}
	return string(result)
}
