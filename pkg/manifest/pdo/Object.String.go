package pdo

import "gopkg.in/yaml.v3"

// String - return a string representation of a parameter dataContract object
func (p Object) String() string {

	object, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(object)

}
