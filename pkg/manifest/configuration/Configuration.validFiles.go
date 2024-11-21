package configuration

// validFiles - validate the configuration file specifications
func (c *Configuration) validFiles() error {
	for _, element := range c.Files {
		if err := element.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}
