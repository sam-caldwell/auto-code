package manifest

import "gopkg.in/yaml.v3"

// String - return a string representation of the HttpResponseMap
func (h *HttpResponseMap) String() string {
	result, err := yaml.Marshal(h)
	if err != nil {
		panic(err)
	}
	return string(result)
}
