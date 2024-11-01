package manifest

import (
	"fmt"
	"regexp"
)

// Verify - validate a given Filename is properly formatted.
func (name FileNameString) Verify() error {

	if pattern := regexp.MustCompile(fileNameStringPattern); pattern.MatchString(string(name)) {

		return nil

	}

	return fmt.Errorf(errInvalidFileName, name)

}
