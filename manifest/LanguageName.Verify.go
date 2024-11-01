package manifest

import (
	"fmt"
	"regexp"
)

// Verify - verify the value is a supported language
func (language LanguageName) Verify() error {
	if pattern := regexp.MustCompile(globalLanguagePattern); pattern.MatchString(string(language)) {
		return nil
	}
	return fmt.Errorf(errInvalidLanguagePattern)
}
