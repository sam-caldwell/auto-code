package manifest

import (
	"fmt"
	"strings"
)

// VerifyType - Verify the property type to ensure it is supported.
func (property *ConfigProperty) VerifyType(name *string) error {

	for _, pType := range strings.Split(configPropertyTypes, Comma) {

		property.Type = strings.ToLower(property.Type)

		if pType == property.Type {
			return nil
		}

	}

	return fmt.Errorf(errUnsupportedPropertyType, *name)

}
