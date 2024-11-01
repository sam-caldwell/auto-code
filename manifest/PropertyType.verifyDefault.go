package manifest

import (
	"fmt"
	"reflect"
)

// verifyDefault - verify that the ConfigProperty.DefaultValue is correct or set its empty state.
func (propertyType *PropertyType) verifyDefault(name *NameIdentifier, value *DefaultValueType) error {

	// If there is no default value we will create the empty-value default for the given type.
	if *value == nil {
		switch string(*propertyType) {
		case "bool":
			*value = false
		case "int":
			*value = int(0)
		case "int8":
			*value = int8(0)
		case "int16":
			*value = int16(0)
		case "int32":
			*value = int32(0)
		case "int64":
			*value = int64(0)
		case "uint":
			*value = uint(0)
		case "uint8":
			*value = uint8(0)
		case "uint16":
			*value = uint16(0)
		case "uint32":
			*value = uint32(0)
		case "uint64":
			*value = uint64(0)
		case "float32":
			*value = float32(0)
		case "float64":
			*value = float64(0)
		case "string":
			*value = EmptyString
		default:
			*value = nil
		}
		return nil
	}

	// if the default value is provided, make sure the type matches the property.type
	if reflect.TypeOf(*value).String() != string(*propertyType) {

		return fmt.Errorf(errDefaultValueTypeMismatch, *name)

	}

	return nil

}
