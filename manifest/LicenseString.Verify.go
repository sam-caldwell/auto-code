package manifest

import (
	"fmt"
	"regexp"
)

// Verify - ensure the value is a valid license string
func (license LicenseString) Verify() error {

	if pattern := regexp.MustCompile(globalLicensePattern); pattern.MatchString(string(license)) {

		return nil

	}

	return fmt.Errorf(errInvalidLicensePattern, globalLicensePattern)

}
