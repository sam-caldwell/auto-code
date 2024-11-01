// Package manifest generator/Manifest.go
package manifest

// ConfigCommandline maps CLI arguments to properties
type ConfigCommandline struct {
	Property string `yaml:"property"`
	Required bool   `yaml:"required"`
	Short    string `yaml:"short"`
	Long     string `yaml:"long"`
}
