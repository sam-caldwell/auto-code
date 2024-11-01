package manifest

// Manifest represents the YAML manifest structure for code generation
type Manifest struct {
	Global struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Author      string `yaml:"author"`
		Email       string `yaml:"email"`
		Copyright   string `yaml:"copyright"`
		License     string `yaml:"license"`
		Language    string `yaml:"language"`
		Version     string `yaml:"version"`
	} `yaml:"global"`

	Config struct {
		Sources     SourceList          `yaml:"sources"`
		Properties  ConfigPropertiesMap `yaml:"properties"`
		File        ConfigFile          `yaml:"file"`
		Environment ConfigEnvironment   `yaml:"environment"`
		Commandline CommandlineArgList  `yaml:"commandline"`
	} `yaml:"config"`
}
