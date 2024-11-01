package manifest

import (
	"fmt"
)

// Verify - verify that config.properties is properly formatted
func (properties ConfigProperties) Verify(_ *ConfigProperties) error {

	if len(properties) < 1 {

		return fmt.Errorf(errMissingConfigProperties)

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
