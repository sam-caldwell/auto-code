// Package manifest generator/Manifest.go
package manifest

import (
	"gopkg.in/yaml.v3"
	"os"
)

// LoadManifest loads and parses the YAML manifest file
func LoadManifest(filePath string) (*Manifest, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var manifest Manifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}
	return &manifest, nil
}
