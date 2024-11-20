package configuration

import "gopkg.in/yaml.v3"

// String - return a string representation of the File struct
func (cfg *File) String() string {
	result, err := yaml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	return string(result)
}
