// Package manifest generator/Manifest.go
package manifest

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Load - loads and parses the YAML manifest file and verifies its content.
func (manifest *Manifest) Load(filePath string) error {

	data, err := os.ReadFile(filePath)

	if err != nil {

		return err

	}

	if err := yaml.Unmarshal(data, manifest); err != nil {

		return err

	}

	if err := manifest.Verify(); err != nil {

		return err

	}

	return nil
}
