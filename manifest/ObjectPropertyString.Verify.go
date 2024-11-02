package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
)

// Verify - Validate the ObjectPropertyString
func (objectPropertyString *ObjectPropertyString) Verify() error {

	if pattern := regexp.MustCompile(patterns.ObjectPropertyStringPattern); pattern.MatchString(string(*objectPropertyString)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidObjectPropertyString, *objectPropertyString)

}
