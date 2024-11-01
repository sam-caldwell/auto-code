package manifest

import (
	"fmt"
	"reflect"
)

// VerifyDefault - verify that the ConfigProperty.DefaultValue is correct or set its empty state.
func (propertyType *PropertyType) VerifyDefault(name *PropertyName, value *DefaultValueType) error {

	// If there is no default value we will create the empty-value default for the given type.
	if *value == nil {

		*value = propertyType.createDefault()
		return nil

	}

	// if the default value is provided, make sure the type matches the property.type
	if reflect.TypeOf(*value).String() != string(*propertyType) {

		return fmt.Errorf(errDefaultValueTypeMismatch, *name)

	}

	return nil

}
