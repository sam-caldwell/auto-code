package data

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/patterns"
	"regexp"
)

// Verify - verify the value is a supported language
func (language LanguageName) Verify() error {
	if pattern := regexp.MustCompile(patterns.LanguagePattern); pattern.MatchString(string(language)) {
		return nil
	}
	return fmt.Errorf(messages.ErrInvalidLanguagePattern)
}
