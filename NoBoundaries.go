package arguments

type NoBoundaries[I int64, U uint64, F float64, S String, B bool] struct{}

func (n NoBoundaries[I, U, F, S, B]) Validator(value any) bool {
	switch value.(type) {
	case int64:
		return true
	case uint64:
		return true
	case float64:
		return true
	case bool:
		return true
	default:
		panic(nonNumericTypeCannotPerformNumericBoundsCheck)
	}
	return false
}
