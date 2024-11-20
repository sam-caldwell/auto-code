package configuration

import "gopkg.in/yaml.v3"

// String - return a string representation of the EnvironmentVariable struct
func (cfg *EnvironmentVariable) String() string {
	result, err := yaml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	return string(result)
}
