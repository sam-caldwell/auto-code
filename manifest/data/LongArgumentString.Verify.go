package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
)

// Verify - validate the long argument string
func (arg *LongArgumentString) Verify() error {

	if *arg == manifest.EmptyString {
		return nil
	}

	if pattern := regexp.MustCompile(patterns.LongArgumentPattern); pattern.MatchString(string(*arg)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidLongArgument, *arg)

}
