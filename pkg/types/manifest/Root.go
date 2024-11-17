package manifest

// Root - represents the top-level of the auto-code YAML specification
type Root struct {
	Version       SemanticVersion `yaml:"auto-code"`
	Info          Info            `yaml:"info"`
	Artifacts     Artifacts       `yaml:"artifacts"`
	Configuration Configuration   `yaml:"configuration"`
	HttpContract  HttpContract    `yaml:"http-contract"`
	DataContract  DataContract    `yaml:"data-contract"`
}
