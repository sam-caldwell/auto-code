package manifest

import (
	"fmt"
	"regexp"
)

// verifyConfigSources - verify config.sources
//
// To add a new supported config.sources, update configSourcesPattern constants
func (manifest *Manifest) verifyConfigSources() error {

	uniqueSources := map[string]struct{}{}

	for _, source := range manifest.Config.Sources {

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
