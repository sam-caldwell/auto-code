package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/patterns"
	"regexp"
)

// verifyGlobalLicense - verify global.license from manifest.yaml
func (manifest *Manifest) verifyGlobalLicense() error {

	if pattern := regexp.MustCompile(patterns.LicensePattern); pattern.MatchString(string(manifest.Global.Name)) {
		return nil
	}

	return fmt.Errorf(messages.ErrInvalidLicensePattern, patterns.LicensePattern)

}
