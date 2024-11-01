package manifest

import (
	"fmt"
	"regexp"
)

// Verify - verify the NameIdentifier to ensure it meets pattern requirements
func (name NameIdentifier) Verify() error {

	if pattern := regexp.MustCompile(NameIdentifierPattern); pattern.MatchString(string(name)) {

		return nil

	}

	return fmt.Errorf(errInvalidNameIdentifier, NameIdentifierPattern)

}
