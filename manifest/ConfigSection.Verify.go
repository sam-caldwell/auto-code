package manifest

// Verify - verify the config section of a manifest.yaml file
func (config *ConfigSection) Verify() error {
	vTable := []func(properties *ConfigProperties) error{
		// We need to verify properties first because other things will use it.
		config.Properties.Verify,
		// Sources must be verified next since it is a potential dependency.
		config.Sources.Verify,
		// All other config properties are data sources and can be in any order.
		config.Commandline.Verify,
		config.Environment.Verify,
		config.File.Verify,

		//ToDo: If you add anything to the config section, add a function to this table to verify its content.

	}
	for _, f := range vTable {
		if err := f(&config.Properties); err != nil {
			return err
		}
	}
	return nil
}
