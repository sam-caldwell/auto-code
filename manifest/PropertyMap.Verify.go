package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
)

// Verify - Validate that the propertyMap key exists in config.Properties
func (propertyMap PropertyMap) Verify(properties *ConfigProperties) error {

	for propertyName, objectPropertyString := range propertyMap {

		// verify our property exists in the set of config.properties
		if _, exists := (*properties)[propertyName]; !exists {
			return fmt.Errorf(messages.ErrUnknownProperty, propertyName)
		}

		// verify our property string is properly formatted
		if err := objectPropertyString.Verify(); err != nil {
			return err
		}

	}

	return nil

}
