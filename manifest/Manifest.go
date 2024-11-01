package manifest

// Manifest represents the YAML manifest structure for code generation
type Manifest struct {

	// Global - defines the Global section representing general program parameters
	Global GlobalSection `yaml:"global"`

	// Config - defines the Config section representing the configuration feature
	Config ConfigSection `yaml:"config"`
}
