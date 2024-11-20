package dataContract

import "gopkg.in/yaml.v3"

// String - return a string representation of the object
func (d *SchemaDescriptorBase) String() string {

	result, err := yaml.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(result)

}
