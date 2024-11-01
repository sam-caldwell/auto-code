package manifest

import (
	"fmt"
	"regexp"
)

// Verify - verify the semantic version string
func (version SemanticVersionString) Verify() error {

	if pattern := regexp.MustCompile(globalVersionPattern); pattern.MatchString(string(version)) {

		return nil

	}

	return fmt.Errorf(errInvalidVersionPattern, globalVersionPattern)

}
