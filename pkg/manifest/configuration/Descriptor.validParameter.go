package configuration

// validParameter - validate that the parameter values are correct.
func (c *Descriptor) validParameter() error {
	for _, element := range c.Parameters {
		if err := element.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}
