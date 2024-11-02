package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"strings"
)

// Verify - verify that a NonEmptyText object contains non-empty text (after stripping whitespace)
func (text NonEmptyText) Verify() error {

	if s := strings.TrimSpace(string(text)); s != manifest.EmptyString {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidNonEmptyText)

}