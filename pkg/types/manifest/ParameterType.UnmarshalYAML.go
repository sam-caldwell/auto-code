package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal a YAML node representing a ParameterType
func (p *ParameterType) UnmarshalYAML(node *yaml.Node) (err error) {
	var rawValue string

	if err := node.Decode(&rawValue); err != nil {
		return err
	}
	switch parts := strings.Split(strings.ToLower(strings.TrimSpace(rawValue)), "("); parts[0] {
	case "string":
		p.base = String
	case "int", "integer":
		p.base = Integer
	case "uint":
		p.base = Uint
	case "uint8":
		p.base = Uint8
	case "uint16":
		p.base = Uint16
	case "uint32":
		p.base = Uint32
	case "uint64":
		p.base = Uint64
	case "float32":
		p.base = Float32
	case "float64":
		p.base = Float64
	case "bool", "boolean":
		p.base = Boolean
	case "object":
		p.base = Object
	case "array":
		p.base = Array
		subParts := strings.Split(strings.TrimSuffix(parts[1], ")"), ",")
		if len(subParts) != 1 {
			return fmt.Errorf("invalid array parameter: '%s'.  Expect 'array(type)'", parts[1])
		}
		switch subParts[0] {
		case "string":
			p.subType = append(p.subType, String)
		case "int", "integer":
			p.subType = append(p.subType, Integer)
		case "uint":
			p.subType = append(p.subType, Uint)
		case "uint8":
			p.subType = append(p.subType, Uint8)
		case "uint16":
			p.subType = append(p.subType, Uint16)
		case "uint32":
			p.subType = append(p.subType, Uint32)
		case "uint64":
			p.subType = append(p.subType, Uint64)
		case "float32":
			p.subType = append(p.subType, Float32)
		case "float64":
			p.subType = append(p.subType, Float64)
		case "bool", "boolean":
			p.subType = append(p.subType, Boolean)
		case "object":
			p.subType = append(p.subType, Object)
		default:
			return fmt.Errorf("invalid array parameter: '%s'.  Expect 'array(type)'", subParts[0])
		}
	case "enum":
		p.base = Enum
		subParts := strings.Split(strings.TrimSuffix(parts[1], ")"), ",")
		if len(subParts) < 1 {
			return fmt.Errorf("invalid array parameter: '%s'.  Expect at least one string", parts[1])
		}
		for _, enumValue := range subParts {
			p.subType = append(p.subType, Identifier) //Enum values are undefined identifiers
			p.metadata = append(p.metadata, enumValue)
		}
	default:

		return fmt.Errorf("unknown/unexpected MergeOrder value (%s)", parts[0])
	}
	return nil
}
