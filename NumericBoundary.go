package arguments

type NumericBoundary[I int64, U uint64, F float64, S string, B bool] struct {
	min any
	max any
}

func (n NumericBoundary[I, U, F, S, B]) Validator(value any) bool {
	if value != nil {
		switch value.(type) {
		case int64:
			return value.(I) >= n.min.(I) && value.(I) <= n.max.(I)
		case uint64:
			return value.(U) >= n.min.(U) && value.(U) <= n.max.(U)
		case float64:
			return value.(F) >= n.min.(F) && value.(F) <= n.max.(F)
		case bool:
			return value.(B) == value.(B)
		default:
			panic(nonNumericTypeCannotPerformNumericBoundsCheck)
		}
		return false
	}
	panic(nilValuePassedToNumericBoundary)
}
