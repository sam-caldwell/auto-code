package manifest

import (
	"fmt"
	"regexp"
)

// VerifyConfigProperties - verify that config.properties is properly formatted
func (manifest *Manifest) verifyConfigProperties() error {

	if len(manifest.Config.Properties) < 1 {
		return fmt.Errorf(errMissingConfigProperties)
	}
	for propertyName, property := range manifest.Config.Properties {

		if pattern := regexp.MustCompile(configPropertyNamePattern); !pattern.MatchString(propertyName) {
			return fmt.Errorf(errInvalidConfigPropertyName)
		}

		if err := manifest.verifyConfigProperty(&propertyName, &property); err != nil {
			return err
		}

	}

	return nil

}
