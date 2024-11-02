package manifest

import (
	"github.com/sam-caldwell/auto-code/validators"
)

// PropertyValidator - validate the config.properties[name].validator
type PropertyValidator struct {
	Class string `yaml:"class"`

	Parameter validator.Interface `yaml:"parameter,omitempty"`
}
