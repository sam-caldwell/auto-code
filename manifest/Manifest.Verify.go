// Package manifest generator/Manifest.go
package manifest

// Verify - Verify the top-level of the manifest.yaml configuration file
func (manifest *Manifest) Verify() error {

	if err := manifest.Global.Verify(); err != nil {
		return err
	}

	if err := manifest.Config.Verify(); err != nil {
		return err
	}

	//
	// ToDo: Extend with other top-level manifest.yaml sections
	//

	return nil

}
