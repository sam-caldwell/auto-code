package manifest

import "fmt"

// Verify - Validate that the property validator object is appropriate for its associated property.
func (validator PropertyValidator) Verify(name *PropertyName, property *ConfigProperty) error {

	switch class := validator.Class; class {

	case "minmax":

		if property.isPropertyTypeNumeric() {

			return verifyParameterMinMax(name, property)

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
