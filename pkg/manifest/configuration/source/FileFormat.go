package source

// FileFormat - The configuration file format (which should match the extension and content)
type FileFormat uint8

const (
	Undefined FileFormat = iota
	YAML
	JSON
)
