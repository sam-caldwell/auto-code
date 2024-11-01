package manifest

import "fmt"

// verifyParameterMinMax - verify that the min/max parameter values are correct
func verifyParameterMinMax(name *PropertyName, property *ConfigProperty) error {

	switch property.Validator.Parameter.(type) {
	case ParameterMinMax[int]:
		return (property.Validator.Parameter.(ParameterMinMax[int])).Verify()
	case ParameterMinMax[int8]:
		return (property.Validator.Parameter.(ParameterMinMax[int8])).Verify()
	case ParameterMinMax[int16]:
		return (property.Validator.Parameter.(ParameterMinMax[int16])).Verify()
	case ParameterMinMax[int32]:
		return (property.Validator.Parameter.(ParameterMinMax[int32])).Verify()
	case ParameterMinMax[int64]:
		return (property.Validator.Parameter.(ParameterMinMax[int64])).Verify()
	case ParameterMinMax[uint]:
		return (property.Validator.Parameter.(ParameterMinMax[uint])).Verify()
	case ParameterMinMax[uint8]:
		return (property.Validator.Parameter.(ParameterMinMax[uint8])).Verify()
	case ParameterMinMax[uint16]:
		return (property.Validator.Parameter.(ParameterMinMax[uint16])).Verify()
	case ParameterMinMax[uint32]:
		return (property.Validator.Parameter.(ParameterMinMax[uint32])).Verify()
	case ParameterMinMax[uint64]:
		return (property.Validator.Parameter.(ParameterMinMax[uint64])).Verify()
	case ParameterMinMax[float32]:
		return (property.Validator.Parameter.(ParameterMinMax[float32])).Verify()
	case ParameterMinMax[float64]:
		return (property.Validator.Parameter.(ParameterMinMax[float64])).Verify()
	default:
		return fmt.Errorf(errExpectedEmptyValidatorParameter, *name)
	}
}
