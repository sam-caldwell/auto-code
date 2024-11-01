package manifest

import (
	"fmt"
	"regexp"
)

// verifyRegex - verify the validator regex parameter is proper.
func (validator ConfigPropertyValidator) verifyRegex(name *PropertyName) (err error) {

	switch validator.Parameter.(type) {
	case string:

		regularExpression := validator.Parameter.(string)

		if _, err = regexp.Compile(regularExpression); err != nil {

			err = fmt.Errorf(errInvalidValidatorRegex, *name, regularExpression, err)

		}

	default:

		err = fmt.Errorf(errExpectedRegexString, *name)

	}

	return err

}
