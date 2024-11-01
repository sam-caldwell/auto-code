package manifest

// createDefault - return a default value of the correct type.
func (propertyType *PropertyType) createDefault() any {

	switch string(*propertyType) {

	case "bool":
		return false

	case "int":
		return int(0)

	case "int8":
		return int8(0)

	case "int16":
		return int16(0)

	case "int32":
		return int32(0)

	case "int64":
		return int64(0)

	case "uint":
		return uint(0)

	case "uint8":
		return uint8(0)

	case "uint16":
		return uint16(0)

	case "uint32":
		return uint32(0)

	case "uint64":
		return uint64(0)

	case "float32":
		return float32(0)

	case "float64":
		return float64(0)

	case "string":

		return EmptyString

	default:
		return nil

	}

}
