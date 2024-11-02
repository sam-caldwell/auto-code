package manifest

// ParameterInterface - is used to allow omitempty behavior for different types of parameters.
//
// Could be a nil, string, []string or ParameterMinMax struct.
type ParameterInterface interface {
	Verify() error
}

type Pattern struct {
	Regex string `yaml:"regex"`
	Match bool   `yaml:"match"`
}

func (p *Pattern) Verify() error { return nil }
