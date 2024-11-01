package manifest

import (
	"fmt"
	"regexp"
)

// verifyGlobalVersion - verify global.version matches the semantic version pattern
func (manifest *Manifest) verifyGlobalVersion() error {
	if pattern := regexp.MustCompile(globalVersionPattern); pattern.MatchString(manifest.Global.Name) {
		return nil
	}
	return fmt.Errorf(errInvalidVersionPattern, globalVersionPattern)
}
