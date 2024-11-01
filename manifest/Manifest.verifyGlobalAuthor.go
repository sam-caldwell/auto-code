package manifest

import (
	"fmt"
	"strings"
)

// verifyGlobalAuthor - verify global.author is valid
func (manifest *Manifest) verifyGlobalAuthor() error {
	if s := strings.TrimSpace(manifest.Global.Author); s != EmptyString {
		return nil
	}
	return fmt.Errorf(errInvalidGlobalAuthor)
}
