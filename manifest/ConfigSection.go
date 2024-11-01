package manifest

// ConfigSection - The Manifest.yaml Config section
type ConfigSection struct {
	Sources SourceList `yaml:"sources"`

	Properties ConfigProperties `yaml:"properties"`

	File ConfigFile `yaml:"file"`

	Environment ConfigEnvironment `yaml:"environment"`

	Commandline CommandlineArgList `yaml:"commandline"`
}
