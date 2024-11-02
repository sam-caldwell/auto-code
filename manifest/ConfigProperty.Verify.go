package manifest

import (
	"github.com/sam-caldwell/auto-code/data"
)

// Verify - verify a given element of config.properties[]
func (property *ConfigProperty) Verify(name *data.NameIdentifier) error {

	if err := name.Verify(); err != nil {
		return err
	}

	if err := property.Type.Verify(name); err != nil {

		return err

	}

	if err := property.Validator.Verify(name, property); err != nil {

		return err

	}

	if err := property.Type.verifyDefault(name, &property.Default); err != nil {

		return err

	}

	return nil

}
