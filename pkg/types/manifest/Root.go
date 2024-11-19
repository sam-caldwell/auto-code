package manifest

// Root - represents the top-level of the auto-code YAML specification
//
// Note: The document root does not support $ref because this would be
//
//	confusing.
type Root struct {
	Version       SemanticVersion `yaml:"auto-code"`
	Info          Info            `yaml:"info"`
	Artifacts     Artifacts       `yaml:"artifacts"`
	Configuration Configuration   `yaml:"configuration"`
	HttpContract  HttpContract    `yaml:"httpContract"`
	DataContract  DataContract    `yaml:"dataContract"`
}
