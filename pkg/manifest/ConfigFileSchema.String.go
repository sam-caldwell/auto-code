package manifest

import "gopkg.in/yaml.v3"

// String - return a string representation of ConfigFileSchema
func (cfg *ConfigFileSchema) String() string {

	result, err := yaml.Marshal(cfg)

	if err != nil {
		panic(err)
	}

	return string(result)

}
