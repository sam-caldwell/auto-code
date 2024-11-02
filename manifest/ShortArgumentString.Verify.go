package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
)

// Verify - verify the short argument string
func (arg *ShortArgumentString) Verify() error {

	if *arg == EmptyString {
		return nil
	}

	if pattern := regexp.MustCompile(patterns.ShortArgumentPattern); pattern.MatchString(string(*arg)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidShortArgument, *arg)

}
