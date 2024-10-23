package arguments

import (
	"fmt"
	"regexp"
)

func (arg *Arguments) String(flag string, defaultValue string, help string, validationRegEx string) *Arguments {

	if pattern := regexp.MustCompile(validationRegEx); pattern.MatchString(defaultValue) {
		arg.args = append(
			arg.args,
			ArgumentDescriptor[int64, uint64, float64, string, bool]{
				flag:  flag,
				value: defaultValue,
				help:  help,
				bounds: PatternBoundary[int64, uint64, float64, string, bool]{
					pattern: pattern,
				},
			})

	} else {
		arg.err = fmt.Errorf(defaultValueMustPassValidation, flag)
	}

	return arg

}
