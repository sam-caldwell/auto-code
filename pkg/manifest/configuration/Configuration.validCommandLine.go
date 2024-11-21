package configuration

// validCommandLine - validate the commandline argument strings
func (c *Configuration) validCommandLine() error {
	for _, element := range c.CommandLine {
		if err := element.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}
