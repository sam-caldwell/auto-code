package manifest

import (
	"fmt"
	"regexp"
)

// Verify - Validate that the environment variable definitions are properly stated.
func (e EnvironmentVariableName) Verify() error {

	if pattern := regexp.MustCompile(environmentVariableNamePattern); !pattern.MatchString(string(e)) {

		return nil

	}

	return fmt.Errorf(errInvalidEnvironmentVariableName, e)

}
