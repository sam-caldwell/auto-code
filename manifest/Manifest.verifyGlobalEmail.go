// Package manifest generator/Manifest.go
package manifest

import (
	"fmt"
	"regexp"
)

func (manifest *Manifest) verifyGlobalEmail() error {
	if pattern := regexp.MustCompile(globalEmailPattern); pattern.MatchString(manifest.Global.Name) {
		return nil
	}
	return fmt.Errorf(errInvalidGlobalEmail)
}
