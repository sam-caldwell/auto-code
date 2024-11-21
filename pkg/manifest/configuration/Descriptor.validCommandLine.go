package configuration

import "fmt"

// validCommandLine - validate the commandline argument strings
func (c *Descriptor) validCommandLine() error {

	for _, element := range c.CommandLine {

		if !c.existsInParameters(element.Parameter.String()) {
			return fmt.Errorf("parameter %s does not exist", element.Name.String())
		}

		if err := element.Name.Validate(); err != nil {
			return err
		}

	}

	return nil

}
