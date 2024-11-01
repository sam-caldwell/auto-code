package manifest

// ConfigPropertyValidator - validate the config.properties[name].validator
type ConfigPropertyValidator struct {
	Class string `yaml:"class"`
	ParameterWrapper
}
