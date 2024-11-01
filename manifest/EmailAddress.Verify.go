package manifest

import (
	"fmt"
	"regexp"
)

// Verify - Validate that the value is a properly formatted email address
func (email EmailAddress) Verify() error {

	if pattern := regexp.MustCompile(globalEmailPattern); pattern.MatchString(string(email)) {

		return nil

	}

	return fmt.Errorf(errInvalidEmailAddress)

}
