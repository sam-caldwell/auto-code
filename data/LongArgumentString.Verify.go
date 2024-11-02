package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/words"
	"regexp"
)

// Verify - validate the long argument string
func (arg *LongArgumentString) Verify() error {

	if *arg == words.EmptyString {
		return nil
	}

	if pattern := regexp.MustCompile(patterns.LongArgumentPattern); pattern.MatchString(string(*arg)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidLongArgument, *arg)

}
