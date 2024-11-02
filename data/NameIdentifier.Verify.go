package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/patterns"
	"regexp"
)

// Verify - verify the NameIdentifier to ensure it meets pattern requirements
func (name NameIdentifier) Verify() error {

	if pattern := regexp.MustCompile(patterns.NameIdentifierPattern); pattern.MatchString(string(name)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidNameIdentifier, patterns.NameIdentifierPattern)

}
