package manifest

import (
	"fmt"
	"regexp"
)

// Verify - verify the short argument string
func (arg *ShortArgumentString) Verify() error {

	if pattern := regexp.MustCompile(shortArgumentPattern); pattern.MatchString(string(*arg)) {

		return nil

	}

	return fmt.Errorf(errInvalidShortArgument, *arg)

}
