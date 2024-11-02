package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/words"
	"regexp"
)

// Verify - verify the short argument string
func (arg *ShortArgumentString) Verify() error {

	if *arg == words.EmptyString {
		return nil
	}

	if pattern := regexp.MustCompile(patterns.ShortArgumentPattern); pattern.MatchString(string(*arg)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidShortArgument, *arg)

}
