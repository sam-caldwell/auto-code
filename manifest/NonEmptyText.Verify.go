package manifest

import (
	"fmt"
	"strings"
)

// Verify - verify that a NonEmptyText object contains non-empty text (after stripping whitespace)
func (text NonEmptyText) Verify() error {

	if s := strings.TrimSpace(string(text)); s != EmptyString {

		return nil

	}

	return fmt.Errorf(errInvalidNonEmptyText)

}
