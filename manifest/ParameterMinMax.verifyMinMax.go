package manifest

import "fmt"

// verifyMinMax - verify that the min/max parameter values are correct
func verifyMinMax(name *string, property *ConfigProperty) error {

	switch property.Validator.Parameter.(type) {
	case ParameterMinMax[int]:
		return (property.Validator.Parameter.(ParameterMinMax[int])).Validate(&property.Type)
	case ParameterMinMax[int8]:
		return (property.Validator.Parameter.(ParameterMinMax[int8])).Validate(&property.Type)
	case ParameterMinMax[int16]:
		return (property.Validator.Parameter.(ParameterMinMax[int16])).Validate(&property.Type)
	case ParameterMinMax[int32]:
		return (property.Validator.Parameter.(ParameterMinMax[int32])).Validate(&property.Type)
	case ParameterMinMax[int64]:
		return (property.Validator.Parameter.(ParameterMinMax[int64])).Validate(&property.Type)
	case ParameterMinMax[uint]:
		return (property.Validator.Parameter.(ParameterMinMax[uint])).Validate(&property.Type)
	case ParameterMinMax[uint8]:
		return (property.Validator.Parameter.(ParameterMinMax[uint8])).Validate(&property.Type)
	case ParameterMinMax[uint16]:
		return (property.Validator.Parameter.(ParameterMinMax[uint16])).Validate(&property.Type)
	case ParameterMinMax[uint32]:
		return (property.Validator.Parameter.(ParameterMinMax[uint32])).Validate(&property.Type)
	case ParameterMinMax[uint64]:
		return (property.Validator.Parameter.(ParameterMinMax[uint64])).Validate(&property.Type)
	case ParameterMinMax[float32]:
		return (property.Validator.Parameter.(ParameterMinMax[float32])).Validate(&property.Type)
	case ParameterMinMax[float64]:
		return (property.Validator.Parameter.(ParameterMinMax[float64])).Validate(&property.Type)
	default:
		return fmt.Errorf(errExpectedEmptyValidatorParameter, *name)
	}
}
