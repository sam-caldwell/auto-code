package source

import "gopkg.in/yaml.v3"

// String - return a string representation of the Environment struct
func (cfg *Environment) String() string {
	result, err := yaml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	return string(result)
}