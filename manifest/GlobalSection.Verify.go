package manifest

// Verify - Verify the global properties section of a manifest.yaml file
func (g *GlobalSection) Verify() error {
	vTable := []func() error{
		g.verifyGlobalName,
		g.verifyGlobalDescription,
		g.verifyGlobalAuthor,
		g.verifyGlobalEmail,
		g.verifyGlobalCopyright,
		g.verifyGlobalLicense,
		g.verifyGlobalLanguage,
		g.verifyGlobalVersion,
		//ToDo: If you add anything to the global section, add a function to this table to verify its content.
	}
	for _, f := range vTable {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
