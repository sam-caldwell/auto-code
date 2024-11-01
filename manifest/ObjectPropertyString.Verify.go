package manifest

import (
	"fmt"
	"regexp"
)

// Verify - Validate the ObjectPropertyString
func (objectPropertyString *ObjectPropertyString) Verify() error {

	if pattern := regexp.MustCompile(objectPropertyStringPattern); pattern.MatchString(string(*objectPropertyString)) {

		return nil

	}

	return fmt.Errorf(errInvalidObjectPropertyString, *objectPropertyString)

}
