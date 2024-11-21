package configuration

import "fmt"

// validFiles - validate the configuration file specifications
func (c *Configuration) validFiles() error {
	for _, element := range c.Files {
		for _, schema := range element.Schema {
			if !c.existsInParameters(schema.Parameter.String()) {
				return fmt.Errorf("parameter %s does not exist", schema.Parameter.String())
			}
		}
		if err := element.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}
