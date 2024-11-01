package manifest

// Manifest represents the YAML manifest structure for code generation
type Manifest struct {
	Global GlobalSection `yaml:"global"`

	Config ConfigSection `yaml:"config"`
}
