package configuration

import "fmt"

// validMergeOrder - ensure that the MergeOrder elements are valid
func (c *Configuration) validMergeOrder() error {
	for _, element := range c.MergeOrder {
		switch element {
		case Files, EnvironmentVars, CommandLine:
			return nil
		default:
			return fmt.Errorf("invalid merge order element: %s", element.String())
		}
	}
	return nil
}
