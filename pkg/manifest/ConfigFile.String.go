package manifest

import "gopkg.in/yaml.v3"

// String - return a string representation of the ConfigFile struct
func (cfg *ConfigFile) String() string {
	result, err := yaml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	return string(result)
}
