package configuration

// FileFormat - The configuration file format (which should match the extension and content)
type FileFormat int

const (
	Undefined FileFormat = iota
	YAML
	JSON
)
