package manifest

// ParameterType - represents a Configuration parameter's data type
type ParameterType struct {
	base     ParameterDataType
	subType  []ParameterDataType
	subValue []any
}

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

func (p *ParameterDataType) String() string {
	return [...]string{
		"string",
		"int",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"boolean",
		"enum",
		"object",
		"array",
	}[*p]
}
