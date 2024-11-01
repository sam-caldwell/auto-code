// Package manifest generator/Manifest.go
package manifest

import (
	"fmt"
	"regexp"
)

// verifyGlobalName - verify the manifest.yaml global.name
func (manifest *Manifest) verifyGlobalName() error {
	if pattern := regexp.MustCompile(globalNamePattern); pattern.MatchString(manifest.Global.Name) {
		return nil
	}
	return fmt.Errorf(errInvalidGlobalName, globalNamePattern)
}
