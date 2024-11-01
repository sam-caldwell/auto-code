package manifest

import (
	"fmt"
	"regexp"
)

// verifyGlobalLanguage - verify global.language from manifest.yaml
func (manifest *Manifest) verifyGlobalLanguage() error {
	if pattern := regexp.MustCompile(globalLanguagePattern); pattern.MatchString(manifest.Global.Name) {
		return nil
	}
	return fmt.Errorf(errInvalidLanguagePattern)
}
