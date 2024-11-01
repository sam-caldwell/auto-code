package manifest

import (
	"fmt"
	"regexp"
)

// verifyRegex - verify the validator regex parameter is proper.
func (v ConfigPropertyValidator) verifyRegex(name *string) (err error) {

	switch v.Parameter.(type) {
	case string:

		regularExpression := v.Parameter.(string)

		if _, err = regexp.Compile(regularExpression); err != nil {

			err = fmt.Errorf(errInvalidValidatorRegex, *name, regularExpression, err)

		}

	default:

		err = fmt.Errorf(errExpectedRegexString, *name)

	}

	return err

}
