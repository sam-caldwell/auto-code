package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/words"
	"strings"
)

// Verify - verify that a NonEmptyText object contains non-empty text (after stripping whitespace)
func (text NonEmptyText) Verify() error {

	if s := strings.TrimSpace(string(text)); s != words.EmptyString {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidNonEmptyText)

}
