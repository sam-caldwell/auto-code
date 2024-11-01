package manifest

// ConfigPropertyValidator - validate the config.properties[name].validator
type ConfigPropertyValidator struct {
	Class string `yaml:"class"`
	ParameterWrapper
}

// ParameterWrapper is used to allow omitempty behavior for different types of parameters.
//
// Could be a nil, string, []string or ParameterMinMax struct.
type ParameterWrapper struct {
	Parameter any `yaml:"parameter,omitempty"`
}
