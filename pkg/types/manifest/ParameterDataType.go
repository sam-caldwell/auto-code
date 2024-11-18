package manifest

// ParameterDataType - represents the data type of ParameterType objects
type ParameterDataType int

const (
	Identifier ParameterDataType = iota
	String
	Integer
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	Boolean
	Enum
	Object
	Array
)
