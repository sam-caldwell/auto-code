package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/patterns"
	"regexp"
)

// Verify - Validate the ObjectPropertyString
func (objectPropertyString *ObjectPropertyString) Verify() error {

	if pattern := regexp.MustCompile(patterns.ObjectPropertyStringPattern); pattern.MatchString(string(*objectPropertyString)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidObjectPropertyString, *objectPropertyString)

}
