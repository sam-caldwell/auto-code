package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
)

// Verify - Validate that the value is a properly formatted email address
func (email EmailAddress) Verify() error {

	if pattern := regexp.MustCompile(patterns.EmailAddressPattern); pattern.MatchString(string(email)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidEmailAddress)

}
