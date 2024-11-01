package manifest

// ParameterWrapper is used to allow omitempty behavior for different types of parameters.
//
// Could be a nil, string, []string or ParameterMinMax struct.
type ParameterWrapper struct {
	Parameter any `yaml:"parameter,omitempty"`
}
