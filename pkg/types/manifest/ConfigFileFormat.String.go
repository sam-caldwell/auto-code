package manifest

// String - return a string representation of a ConfigFileFormat enumerated type
func (cfg *ConfigFileFormat) String() string {
	switch *cfg {
	case YAML:
		return "yaml"
	case JSON:
		return "json"
	default:
		panic("unknown or unsupported Config file format (use YAML/JSON)")
	}
}
