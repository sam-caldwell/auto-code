package manifest

// VerifyConfigProperty - verify a given element of config.properties[]
func (manifest *Manifest) verifyConfigProperty(name *string, property *ConfigProperty) error {

	if err := property.VerifyType(name); err != nil {
		return err
	}

	if err := property.VerifyValidator(name); err != nil {
		return err
	}

	if err := property.verifyDefault(name); err != nil {
		return err
	}

	return nil
}
