package manifest

// ConfigProperty defines each configurable property
type ConfigProperty struct {
	Type      PropertyType      `yaml:"type"`
	Validator PropertyValidator `yaml:"validator"`
	Default   DefaultValueType  `yaml:"default,omitempty"`
}
