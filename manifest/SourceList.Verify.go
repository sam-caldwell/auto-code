package manifest

import (
	"fmt"
	"regexp"
)

// Verify - Verify that the source list is correct (non-empty, containing valid values)
//
// WARNING: At no point can the SourceList order be altered.  That would break merge ordering.
func (s SourceList) Verify() error {

	uniqueSources := map[string]struct{}{}

	for _, source := range s {

		if pattern := regexp.MustCompile(configSourcesPattern); pattern.MatchString(source) {
			return fmt.Errorf(errInvalidConfigSource, configSourcesPattern)
		}

		if _, ok := uniqueSources[source]; ok {
			return fmt.Errorf(errDuplicateConfigSource)
		}

		uniqueSources[source] = struct{}{}

	}

	if numberOfSources := len(uniqueSources); numberOfSources < 1 {
		return fmt.Errorf(errMissingConfigSource)
	}

	return nil

}
