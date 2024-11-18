package manifest

// ParameterType - represents a Configuration parameter's data type
type ParameterType struct {
	base     ParameterDataType
	subType  []ParameterDataType
	metadata []any
}
