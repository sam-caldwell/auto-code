package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
)

// Verify - verify the semantic version string
func (version SemanticVersionString) Verify() error {

	if pattern := regexp.MustCompile(patterns.SemanticVersionPattern); pattern.MatchString(string(version)) {

		return nil

	}

	return fmt.Errorf(messages.ErrInvalidVersionPattern, patterns.SemanticVersionPattern)

}
