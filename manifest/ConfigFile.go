// Package manifest generator/Manifest.go
package manifest

// ConfigFile defines file-based configuration
type ConfigFile struct {
	File     string            `yaml:"file"`
	Required bool              `yaml:"required"`
	Format   string            `yaml:"format"`
	Map      map[string]string `yaml:"map"`
}
