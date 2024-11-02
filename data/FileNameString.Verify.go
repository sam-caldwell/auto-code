package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/patterns"
	"regexp"
)

// Verify - validate a given Filename is properly formatted.
func (name FileNameString) Verify() error {

	if pattern := regexp.MustCompile(patterns.FileNameStringPattern); pattern.MatchString(string(name)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidFileName, name)

}
