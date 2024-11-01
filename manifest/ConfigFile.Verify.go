package manifest

// Verify - Validate the configuration file section of manifest.yaml
func (file ConfigFile) Verify(properties *ConfigProperties) error {

	if err := file.File.Verify(); err != nil {
		return err
	}

	if err := file.Format.Verify(); err != nil {
		return err
	}

	if err := file.Map.Verify(properties); err != nil {
		return err
	}

	return nil

}
