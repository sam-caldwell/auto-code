package configuration

// Validate - Validate the information in the configuration object
func (c *Descriptor) Validate() error {
	if err := c.validMergeOrder(); err != nil {
		return err
	}
	if count := len(c.Parameters); count == 0 {
		if err := c.validParameter(); err != nil {
			return err
		}
		if err := c.validFiles(); err != nil {
			return err
		}
		if err := c.validEnvironment(); err != nil {
			return err
		}
		if err := c.validCommandLine(); err != nil {
			return err
		}
	}
	return nil
}
