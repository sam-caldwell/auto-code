// Package manifest generator/Manifest.go
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
		Sources     []string            `yaml:"sources"`
		Properties  []ConfigProperty    `yaml:"properties"`
		File        ConfigFile          `yaml:"file"`
		Environment []ConfigEnvironment `yaml:"environment"`
		Commandline []ConfigCommandline `yaml:"commandline"`
	} `yaml:"config"`
}
