package archive

import "gopkg.in/yaml.v3"

func (d *HttpEndpointDescriptor) String() string {
	result, err := yaml.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(result)
}
