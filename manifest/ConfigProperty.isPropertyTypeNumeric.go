package manifest

// isNumeric - Return boolean indicating if property.Type is numeric
func (property *ConfigProperty) isPropertyTypeNumeric() bool {
	switch property.Type {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":

		return true

	default:
		return false
	}
}
