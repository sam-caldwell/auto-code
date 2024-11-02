package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/patterns"
	"regexp"
)

// Verify - ensure the value is a valid license string
func (license LicenseString) Verify() error {

	if pattern := regexp.MustCompile(patterns.LicensePattern); pattern.MatchString(string(license)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidLicensePattern, patterns.LicensePattern)

}
