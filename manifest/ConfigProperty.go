package manifest

import (
	"github.com/sam-caldwell/auto-code/data"
)

// ConfigProperty defines each configurable property
type ConfigProperty struct {
	Type      data.PropertyType `yaml:"type"`
	Validator PropertyValidator `yaml:"validator"`
	Default   DefaultValueType  `yaml:"default,omitempty"`
}
