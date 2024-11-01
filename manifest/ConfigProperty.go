package manifest

// ConfigProperty defines each configurable property
type ConfigProperty struct {
	Type      string                  `yaml:"type"`
	Validator ConfigPropertyValidator `yaml:"validator"`
	Default   interface{}             `yaml:"default"`
}
