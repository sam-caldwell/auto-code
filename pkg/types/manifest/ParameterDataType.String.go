package manifest

// String - return a string representation of a parameter data type
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
