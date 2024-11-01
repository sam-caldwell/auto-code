package manifest

import (
	"fmt"
	"regexp"
)

// Verify - Validate that the property validator object is appropriate for its associated property.
func (validator PropertyValidator) Verify(name *NameIdentifier, property *ConfigProperty) (err error) {

	switch class := validator.Class; class {
	/*
	 * Min-Max validators are expected to be numeric property types.
	 */
	case "minmax":
		switch property.Validator.Parameter.(type) {
		case ParameterMinMax[int]:
			err = (property.Validator.Parameter.(ParameterMinMax[int])).Verify()
		case ParameterMinMax[int8]:
			err = (property.Validator.Parameter.(ParameterMinMax[int8])).Verify()
		case ParameterMinMax[int16]:
			err = (property.Validator.Parameter.(ParameterMinMax[int16])).Verify()
		case ParameterMinMax[int32]:
			err = (property.Validator.Parameter.(ParameterMinMax[int32])).Verify()
		case ParameterMinMax[int64]:
			err = (property.Validator.Parameter.(ParameterMinMax[int64])).Verify()
		case ParameterMinMax[uint]:
			err = (property.Validator.Parameter.(ParameterMinMax[uint])).Verify()
		case ParameterMinMax[uint8]:
			err = (property.Validator.Parameter.(ParameterMinMax[uint8])).Verify()
		case ParameterMinMax[uint16]:
			err = (property.Validator.Parameter.(ParameterMinMax[uint16])).Verify()
		case ParameterMinMax[uint32]:
			err = (property.Validator.Parameter.(ParameterMinMax[uint32])).Verify()
		case ParameterMinMax[uint64]:
			err = (property.Validator.Parameter.(ParameterMinMax[uint64])).Verify()
		case ParameterMinMax[float32]:
			err = (property.Validator.Parameter.(ParameterMinMax[float32])).Verify()
		case ParameterMinMax[float64]:
			err = (property.Validator.Parameter.(ParameterMinMax[float64])).Verify()
		default:
			err = fmt.Errorf(errExpectedEmptyValidatorParameter, *name)
		}
	/*
	 * null validators have no validation.
	 */
	case "null":
		if property.Validator.Parameter != nil {
			err = fmt.Errorf(errExpectedEmptyValidatorParameter, *name)
		}
	/*
	 * regular expressions expect a string property type.
	 */
	case "regex":
		switch validator.Parameter.(type) {
		case string:
			regularExpression := validator.Parameter.(string)

			if _, err = regexp.Compile(regularExpression); err != nil {
				err = fmt.Errorf(errInvalidValidatorRegex, *name, regularExpression, err)
			}
		default:
			err = fmt.Errorf(errExpectedRegexString, *name)
		}
	/*
	 * unsupported/unknown validator type
	 */
	default:
		err = fmt.Errorf(errUnsupportedPropertyValidator, *name)
	}
	return err
}
