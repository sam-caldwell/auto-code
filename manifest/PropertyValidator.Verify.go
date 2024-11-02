package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/data"
	"github.com/sam-caldwell/auto-code/messages"
	validator2 "github.com/sam-caldwell/auto-code/validators"
	"regexp"
)

// Verify - Validate that the property validator object is appropriate for its associated property.
func (v PropertyValidator) Verify(name *data.NameIdentifier, property *ConfigProperty) (err error) {

	switch class := v.Class; class {
	/*
	 * Min-Max validators are expected to be numeric property types.
	 */
	case "minmax":
		switch property.Validator.Parameter.(type) {
		case validator2.Int:
			err = (property.Validator.Parameter.(validator2.Int)).Verify()
		case validator2.Int8:
			err = (property.Validator.Parameter.(validator2.Int8)).Verify()
		case validator2.Int16:
			err = (property.Validator.Parameter.(validator2.Int16)).Verify()
		case validator2.Int32:
			err = (property.Validator.Parameter.(validator2.Int32)).Verify()
		case validator2.Int64:
			err = (property.Validator.Parameter.(validator2.Int64)).Verify()
		case validator2.Uint:
			err = (property.Validator.Parameter.(validator2.Uint)).Verify()
		case validator2.Uint8:
			err = (property.Validator.Parameter.(validator2.Uint8)).Verify()
		case validator2.Uint16:
			err = (property.Validator.Parameter.(validator2.Uint16)).Verify()
		case validator2.Uint32:
			err = (property.Validator.Parameter.(validator2.Uint32)).Verify()
		case validator2.Uint64:
			err = (property.Validator.Parameter.(validator2.Uint64)).Verify()
		case validator2.Float32:
			err = (property.Validator.Parameter.(validator2.Float32)).Verify()
		case validator2.Float64:
			err = (property.Validator.Parameter.(validator2.Float64)).Verify()
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
