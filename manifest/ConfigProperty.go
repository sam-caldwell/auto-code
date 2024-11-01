// Package manifest generator/Manifest.go
package manifest

// ConfigProperty defines each configurable property
type ConfigProperty struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Validator struct {
		Class     string `yaml:"class"`
		Parameter string `yaml:"parameter"`
	} `yaml:"validator"`
	Default interface{} `yaml:"default"`
}
