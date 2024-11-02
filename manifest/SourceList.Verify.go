package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
)

// Verify - Verify that the source list is correct (non-empty, containing valid values)
//
// WARNING: At no point can the SourceList order be altered.  That would break merge ordering.
func (s SourceList) Verify(_ *ConfigProperties) error {

	uniqueSources := map[string]struct{}{}

	for _, source := range s {

		if pattern := regexp.MustCompile(patterns.ConfigSourcesPattern); pattern.MatchString(source) {
			return fmt.Errorf(messages.ErrInvalidConfigSource, patterns.ConfigSourcesPattern)
		}

		if _, ok := uniqueSources[source]; ok {
			return fmt.Errorf(messages.ErrDuplicateConfigSource)
		}

		uniqueSources[source] = struct{}{}

	}

	if numberOfSources := len(uniqueSources); numberOfSources < 1 {
		return fmt.Errorf(messages.ErrMissingConfigSource)
	}

	return nil

}
