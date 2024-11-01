package manifest

// Verify - verify the config section of a manifest.yaml file
func (config *ConfigSection) Verify() error {
	vTable := []func() error{
		config.Commandline.Verify,
		config.Environment.Verify,
		config.File.Verify,
		config.Sources.Verify,
		config.Properties.Verify,
		//ToDo: If you add anything to the config section, add a function to this table to verify its content.
		//      the order of verification functions does not matter.  We probably should keep it alphabetical.
	}
	for _, f := range vTable {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
