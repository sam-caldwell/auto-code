package manifest

// ParameterValue - represents a parameter value, type and validation regex
type ParameterValue struct {
	DataObjectWithReference
	Data ParameterDataObject `yaml:"value"`
}
