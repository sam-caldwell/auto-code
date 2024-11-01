package manifest

import (
	"fmt"
	"regexp"
)

// verifyGlobalLicense - verify global.license from manifest.yaml
func (manifest *Manifest) verifyGlobalLicense() error {

	if pattern := regexp.MustCompile(globalLicensePattern); pattern.MatchString(string(manifest.Global.Name)) {
		return nil
	}

	return fmt.Errorf(errInvalidLicensePattern, globalLicensePattern)

}
