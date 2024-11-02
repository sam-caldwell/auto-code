package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	validator "github.com/sam-caldwell/auto-code/manifest/validators"
	"regexp"
)

// Verify - Validate that the property validator object is appropriate for its associated property.
func (v PropertyValidator) Verify(name *NameIdentifier, property *ConfigProperty) (err error) {

	switch class := v.Class; class {
	/*
	 * Min-Max validators are expected to be numeric property types.
	 */
	case "minmax":
		switch property.Validator.Parameter.(type) {
		case validator.Int:
			err = (property.Validator.Parameter.(validator.Int)).Verify()
		case validator.Int8:
			err = (property.Validator.Parameter.(validator.Int8)).Verify()
		case validator.Int16:
			err = (property.Validator.Parameter.(validator.Int16)).Verify()
		case validator.Int32:
			err = (property.Validator.Parameter.(validator.Int32)).Verify()
		case validator.Int64:
			err = (property.Validator.Parameter.(validator.Int64)).Verify()
		case validator.Uint:
			err = (property.Validator.Parameter.(validator.Uint)).Verify()
		case validator.Uint8:
			err = (property.Validator.Parameter.(validator.Uint8)).Verify()
		case validator.Uint16:
			err = (property.Validator.Parameter.(validator.Uint16)).Verify()
		case validator.Uint32:
			err = (property.Validator.Parameter.(validator.Uint32)).Verify()
		case validator.Uint64:
			err = (property.Validator.Parameter.(validator.Uint64)).Verify()
		case validator.Float32:
			err = (property.Validator.Parameter.(validator.Float32)).Verify()
		case validator.Float64:
			err = (property.Validator.Parameter.(validator.Float64)).Verify()
		default:
			err = fmt.Errorf(messages.ErrExpectedEmptyValidatorParameter, *name)
		}
	/*
	 * null validators have no validation.
	 */
	case "null":
		if property.Validator.Parameter != nil {
			err = fmt.Errorf(messages.ErrExpectedEmptyValidatorParameter, *name)
		}
	/*
	 * regular expressions expect a string property type.
	 */
	case "regex":
		switch v.Parameter.(type) {
		case string:
			regularExpression := v.Parameter.(string)

			if _, err = regexp.Compile(regularExpression); err != nil {
				err = fmt.Errorf(messages.ErrInvalidValidatorRegex, *name, regularExpression, err)
			}
		default:
			err = fmt.Errorf(messages.ErrExpectedRegexString, *name)
		}
	/*
	 * unsupported/unknown validator type
	 */
	default:
		err = fmt.Errorf(messages.ErrUnsupportedPropertyValidator, *name)
	}
	return err
}
