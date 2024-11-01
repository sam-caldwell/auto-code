package manifest

import (
	"fmt"
	"strings"
)

// Verify - Validate the PropertyType and ensure it is lowercase and has no leading or trailing whitespace.
func (propertyType *PropertyType) Verify(name *NameIdentifier) error {

	for _, pType := range strings.Split(configPropertyTypes, Comma) {

		// strip the property type and make sure it is lower case.
		*propertyType = PropertyType(strings.ToLower(string(*propertyType)))

		if pType == string(*propertyType) {

			return nil

		}

	}

	return fmt.Errorf(errUnsupportedPropertyType, *name)

}
