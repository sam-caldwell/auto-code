// Package manifest generator/Manifest.go
package manifest

// Commandline - maps CLI arguments to properties
type Commandline struct {
	Property string `yaml:"property"`
	Required bool   `yaml:"required"`
	Short    string `yaml:"short"`
	Long     string `yaml:"long"`
}
