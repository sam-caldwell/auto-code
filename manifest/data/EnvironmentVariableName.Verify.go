package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
)

// Verify - Validate that the environment variable definitions are properly stated.
func (e EnvironmentVariableName) Verify() error {

	if pattern := regexp.MustCompile(patterns.environmentVariableNamePattern); !pattern.MatchString(string(e)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidEnvironmentVariableName, e)

}
