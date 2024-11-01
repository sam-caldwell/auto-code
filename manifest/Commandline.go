// Package manifest generator/Manifest.go
package manifest

// Commandline - maps CLI arguments to properties
type Commandline struct {
	Required bool `yaml:"required"`

	Short ShortArgumentString `yaml:"short,omitempty"`

	Long LongArgumentString `yaml:"long,omitempty"`
}
