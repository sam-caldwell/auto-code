package manifest

import (
	"fmt"
	"regexp"
)

// Verify - Validate the property name is correctly formatted
func (name PropertyName) Verify() error {

	pattern := regexp.MustCompile(configPropertyNamePattern)

	if ok := pattern.MatchString(string(name)); ok {
		return nil
	}

	return fmt.Errorf(errInvalidConfigPropertyName)

}
