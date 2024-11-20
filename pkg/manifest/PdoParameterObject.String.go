package manifest

import "gopkg.in/yaml.v3"

// String - return a string representation of a parameter data object
func (p PdoParameterObject) String() string {
	object, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(object)
}
