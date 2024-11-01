package manifest

// verifyGlobal - verify the manifest.yaml global section
func (manifest *Manifest) verifyGlobal() error {
	vTable := []func() error{
		manifest.verifyGlobalName,
		manifest.verifyGlobalDescription,
		manifest.verifyGlobalAuthor,
		manifest.verifyGlobalEmail,
		manifest.verifyGlobalCopyright,
		manifest.verifyGlobalLicense,
		manifest.verifyGlobalLanguage,
		manifest.verifyGlobalVersion,
		//ToDo: If you add anything to the global section, add a function to this table to verify its content.
	}
	for _, f := range vTable {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
