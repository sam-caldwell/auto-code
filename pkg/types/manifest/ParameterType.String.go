package manifest

import (
	"fmt"
)

// String - returns a string representation of a ParameterType
func (p *ParameterType) String() string {

	var typeStrings = [...]string{
		"string",
		"int",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"boolean",
		"enum",
		"object",
		"array",
	}

	switch p.base {
	case Integer, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, Boolean, Object:

		return typeStrings[p.base]

	case Array:

		return fmt.Sprintf("%s(%s)", typeStrings[p.base], p.subType[0].String())

	case Enum:
		var subValues string
		for _, v := range p.metadata {
			if element, ok := v.(string); ok {
				subValues += element
			} else {
				panic("programming error (enum type improperly stored")
			}
		}
		return fmt.Sprintf("%s(%s)", typeStrings[p.base], subValues)

	default:
		panic("programming error (invalid ParameterType object)")
	}

}
