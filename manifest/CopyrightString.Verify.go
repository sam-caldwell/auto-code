package manifest

import (
	"fmt"
	"regexp"
)

// Verify - Validate that the value is a proper copyright string
func (copyright CopyrightString) Verify() error {

	if pattern := regexp.MustCompile(globalCopyrightPattern); pattern.MatchString(string(copyright)) {

		return nil

	}

	return fmt.Errorf(errInvalidCopyrightPattern)

}
