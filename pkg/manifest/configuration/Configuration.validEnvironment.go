package configuration

// validEnvironment - validate the environment variable names
func (c *Configuration) validEnvironment() error {
	for _, element := range c.Environment {
		if err := element.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}
