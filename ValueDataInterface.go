package arguments

type ValueDataInterface interface {
	Action(value any) bool
	Data() any
	Error() error
}

//// Int - Creates a signed integer value
//func Int(defaultValue int, validator *Validator) *Value {
//	return &Value{
//		data:      defaultValue,
//		validator: validator,
//	}
//}
//
//// Int64 - Creates a signed 64-bit integer value
//func Int64(defaultValue int64, validator *Validator) *Value {
//	return &Value{
//		data:      defaultValue,
//		validator: validator,
//	}
//}
//
//// Float64 - Creates a floating-point value
//func Float64(defaultValue float64, validator *Validator) *Value {
//	return &Value{
//		data:      defaultValue,
//		validator: validator,
//	}
//}
//
//// Uint - Creates an unsigned integer value
//func Uint(defaultValue uint, validator *Validator) *Value {
//	return &Value{
//		data:      defaultValue,
//		validator: validator,
//	}
//}
//
//// Uint64 - Creates an unsigned 64-bit integer value
//func Uint64(defaultValue uint64, validator *Validator) *Value {
//	return &Value{
//		data:      defaultValue,
//		validator: validator,
//	}
//}
