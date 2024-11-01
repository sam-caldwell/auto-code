// Package manifest generator/Manifest.go
package manifest

// ConfigFile defines file-based configuration
type ConfigFile struct {
	File FileNameString `yaml:"file"`

	Required bool `yaml:"required,omitempty"`

	Format FileFormatString `yaml:"format"`

	Map PropertyMap `yaml:"map"`
}

// ObjectPropertyString - a string representing a dot-delimited object property name
type ObjectPropertyString struct {
}
