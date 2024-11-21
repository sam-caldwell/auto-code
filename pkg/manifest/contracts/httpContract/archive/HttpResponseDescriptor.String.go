package archive

import "gopkg.in/yaml.v3"

// String - return a string representation of HttpResponseDescriptor
func (h *HttpResponseDescriptor) String() string {
	result, err := yaml.Marshal(h)
	if err != nil {
		panic(err)
	}
	return string(result)
}
