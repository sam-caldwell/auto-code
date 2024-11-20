package manifest

// ParameterValue - represents a parameter value, type and validation regex
//
// # This object supports external $ref references
//
// This object implements a ParameterDataObject (PDO) representing a strongly
// typed data value (state) and its validator logic.  While the PDO can store
// many different data types, it does perform type-checking and enforcement.
type ParameterValue struct {
	DataObjectWithReference
	Data ParameterDataObject `yaml:"value"`
}
