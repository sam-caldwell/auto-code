package manifest

// PropertyValidator - validate the config.properties[name].validator
type PropertyValidator struct {
	Class string `yaml:"class"`

	ParameterWrapper
}
