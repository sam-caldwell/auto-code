package arguments

import "fmt"

// Type - Create a typed value data object that implements ValueDataInterface
//
// This function accepts any defaultValue and validator then verifies that the
// validator and defaultValue agree and make sense before returning a typed
// data object implementing ValueDataInterface based on the defaultValue type.
func Type(defaultValue any, validator *Validator) ValueDataInterface {
	if validator == nil {
		return InvalidDataObject{
			fmt.Errorf(nilValidatorObject),
		}
	}
	if !validator.action(defaultValue) {
		return InvalidDataObject{
			fmt.Errorf(invalidDefaultValue),
		}
	}

	switch defaultValue.(type) {
	case string:
		return StringDataObject{
			data:      defaultValue.(string),
			validator: validator,
			err:       nil,
		}
	default:
		return InvalidDataObject{
			fmt.Errorf(invalidValueType),
		}
	}
}
