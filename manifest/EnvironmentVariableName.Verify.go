package manifest

import (
	"fmt"
	"regexp"
)

func (e EnvironmentVariableName) Verify() error {
	if pattern := regexp.MustCompile(environmentVariableNamePattern); pattern.MatchString(string(e)) {
		return nil
	}
	return fmt.Errorf(errInvalidEnvironmentVariableName, e)
}
