package arguments

import "regexp"

type PatternBoundary[I int64, U uint64, F float64, S string, B bool] struct {
	pattern *regexp.Regexp
}

func (p PatternBoundary[I, U, F, S, B]) Validator(value any) bool {
	if value != nil {
		switch value.(type) {
		case string:
			return p.pattern.MatchString(value.(string))
		case []byte:
			return p.pattern.MatchString(string(value.([]byte)))
		default:
			panic(numericTypeCannotPerformPatternBoundsCheck)
		}
		return false
	}
	panic(nilValuePassedToNumericBoundary)
}
