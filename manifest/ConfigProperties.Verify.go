package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
)

// Verify - verify that config.properties is properly formatted
func (properties ConfigProperties) Verify(_ *ConfigProperties) error {

	if len(properties) < 1 {

		return fmt.Errorf(messages.ErrMissingConfigProperties)

	}

	for propertyName, property := range properties {

		if err := propertyName.Verify(); err != nil {
			return err
		}

		if err := property.Verify(&propertyName); err != nil {
			return err
		}

	}

	return nil

}
