// Package manifest generator/Manifest.go
package manifest

// ConfigEnvironment maps environment variables to properties
type ConfigEnvironment struct {
	Property string `yaml:"property"`
	EnvVar   string `yaml:"env_var"`
}
