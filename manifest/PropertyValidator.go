package manifest

import validator "github.com/sam-caldwell/auto-code/manifest/validators"

// PropertyValidator - validate the config.properties[name].validator
type PropertyValidator struct {
	Class string `yaml:"class"`

	Parameter validator.Interface `yaml:"parameter,omitempty"`
}
