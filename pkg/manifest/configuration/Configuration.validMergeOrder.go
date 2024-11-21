package configuration

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/pkg/types/generic/enum"
)

// validMergeOrder - ensure that the MergeOrder elements are valid
func (c *Configuration) validMergeOrder() error {
	for _, element := range c.MergeOrder {
		switch element {
		case enum.Files, enum.EnvironmentVars, enum.CommandLine:
			return nil
		default:
			return fmt.Errorf("invalid merge order element: %s", element.String())
		}
	}
	return nil
}
