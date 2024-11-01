package manifest

import (
	"fmt"
	"regexp"
)

// Verify - validate the long argument string
func (arg *LongArgumentString) Verify() error {

	if pattern := regexp.MustCompile(longArgumentPattern); pattern.MatchString(string(*arg)) {

		return nil

	}

	return fmt.Errorf(errInvalidLongArgument, *arg)

}
