package manifest

import (
	"fmt"
	"strings"
)

// verifyGlobalDescription - verify global.descriptor is valid
func (manifest *Manifest) verifyGlobalDescription() error {
	if s := strings.TrimSpace(manifest.Global.Description); s != EmptyString {
		return nil
	}
	return fmt.Errorf(errInvalidGlobalDescription)
}
