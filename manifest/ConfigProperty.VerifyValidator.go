package manifest

import (
	"fmt"
)

// VerifyValidator - Verify the property default value to ensure it is supported.
func (property *ConfigProperty) VerifyValidator(name *string) error {

	switch class := property.Validator.Class; class {
	case "minmax":
		if property.isPropertyTypeNumeric() {
			return verifyMinMax(name, property)
		}
		return fmt.Errorf(errTypeMismatchNumeric, *name)
	case "null":
		if property.Validator.Parameter != nil {
			return fmt.Errorf(errExpectedEmptyValidatorParameter, *name)
		}
	case "regex":
		if property.Type != "string" {
			return fmt.Errorf(errTypeMismatchString, *name)
		}
		return property.Validator.verifyRegex(name)
	}
	return fmt.Errorf(errUnsupportedPropertyValidator, *name)

}
