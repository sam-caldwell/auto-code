package manifest

// ParameterDataObject - represents the actual dataContract value and its associated validation logic.
type ParameterDataObject struct {
	DataObjectWithReference
	State     any              `yaml:"default,omitempty"`
	validator PdoDataValidator `yaml:"regex,omitempty"`
}
