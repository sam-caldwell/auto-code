package enum

// String - return a string representation of a FileFormat enumerated type
func (cfg *FileFormat) String() string {
	switch *cfg {
	case YAML:
		return "yaml"
	case JSON:
		return "json"
	default:
		panic("unknown or unsupported Config file format (use YAML/JSON)")
	}
}
