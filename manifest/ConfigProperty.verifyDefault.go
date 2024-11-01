package manifest

import (
	"fmt"
	"reflect"
)

// VerifyDefault - Verify the property default value type to ensure it matches the rest of the config.
func (property *ConfigProperty) verifyDefault(name *string) error {

	defaultValue := property.Default

	// If there is no default value we will create the empty-value default for the given type.
	if property.Default == nil {
		property.Default = property.createDefault()
		return nil
	}

	// if the default value is provided, make sure the type matches the property.type
	if reflect.TypeOf(defaultValue).String() != property.Type {
		return fmt.Errorf(errDefaultValueTypeMismatch, *name)
	}

	return nil

}
