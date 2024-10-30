package arguments

import (
	"fmt"
	"regexp"
)

type Validator struct {
	action func(v any) bool
}

func Pattern(pattern string) *Validator {
	return &Validator{
		action: func(v any) bool {
			re := regexp.MustCompile(pattern)
			return re.Match([]byte(fmt.Sprintf("%v", v)))
		},
	}
}

func MinMaxInt(min, max int64) *Validator {
	return &Validator{
		action: func(v any) bool {
			if value, ok := v.(int64); ok {
				return min <= value && value <= max
			} else {
				return false
			}
		},
	}
}
func MinMaxFloat(min, max float64) *Validator {
	return &Validator{
		action: func(v any) bool {
			if value, ok := v.(float64); ok {
				return min <= value && value <= max
			} else {
				return false
			}
		},
	}
}
