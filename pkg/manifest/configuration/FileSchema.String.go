package configuration

import "gopkg.in/yaml.v3"

// String - return a string representation of FileSchema
func (cfg *FileSchema) String() string {

	result, err := yaml.Marshal(cfg)

	if err != nil {
		panic(err)
	}

	return string(result)

}
