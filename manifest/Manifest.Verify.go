// Package manifest generator/Manifest.go
package manifest

// Verify - Verify the manifest.yaml file contains valid data.
func (manifest *Manifest) Verify() error {
	if err := manifest.verifyGlobal(); err != nil {
		return err
	}
	if err := manifest.verifyConfig(); err != nil {
		return err
	}
	//
	// ToDo: Extend with other top-level manifest.yaml sections
	//
	return nil
}
