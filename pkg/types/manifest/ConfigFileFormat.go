package manifest

// ConfigFileFormat - The configuration file format (which should match the extension and content)
type ConfigFileFormat int

const (
	Undefined ConfigFileFormat = iota
	YAML
	JSON
)
