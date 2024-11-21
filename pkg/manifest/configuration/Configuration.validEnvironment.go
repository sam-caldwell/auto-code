package configuration

import "fmt"

// validEnvironment - validate the environment variable names
func (c *Configuration) validEnvironment() error {

	for _, element := range c.Environment {

		if !c.existsInParameters(element.Parameter.String()) {
			return fmt.Errorf("parameter %s does not exist", element.Name.String())
		}

		if err := element.Name.Validate(); err != nil {
			return err
		}

	}

	return nil

}
