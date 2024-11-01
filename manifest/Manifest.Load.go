// Package manifest generator/Manifest.go
package manifest

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Load - loads and parses the YAML manifest file and verifies its content.
func (manifest *Manifest) Load(filePath string) (err error) {

	var data []byte

	if data, err = os.ReadFile(filePath); err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, manifest); err != nil {
		return err
	}

	// get the local repo url and merge it with what may have been defined in the yaml file
	if err := manifest.GetLocalRepo(); err != nil {
		return err
	}

	// verify the manifest content
	if err := manifest.Verify(); err != nil {
		return err
	}

	return nil
}
