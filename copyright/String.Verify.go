package copyright

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/patterns"
	"regexp"
)

// Verify - Validate that the value is a proper copyright string
func (copyright String) Verify() error {

	if pattern := regexp.MustCompile(patterns.CopyrightPattern); pattern.MatchString(string(copyright)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidCopyrightPattern)

}
