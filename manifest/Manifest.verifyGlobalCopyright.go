package manifest

import (
	"fmt"
	"regexp"
)

// verifyGlobalCopyright - verify the global.copyright is properly formatted.
func (manifest *Manifest) verifyGlobalCopyright() error {
	if pattern := regexp.MustCompile(globalCopyrightPattern); pattern.MatchString(manifest.Global.Name) {
		return nil
	}
	return fmt.Errorf(errInvalidCopyrightPattern)
}
