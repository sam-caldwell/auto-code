package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
)

// Verify - validate a given Filename is properly formatted.
func (name FileNameString) Verify() error {

	if pattern := regexp.MustCompile(patterns.FileNameStringPattern); pattern.MatchString(string(name)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidFileName, name)

}
