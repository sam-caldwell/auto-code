package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
)

// Verify - Validate that the value is a proper copyright string
func (copyright CopyrightString) Verify() error {

	if pattern := regexp.MustCompile(patterns.CopyrightPattern); pattern.MatchString(string(copyright)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidCopyrightPattern)

}
