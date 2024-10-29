package arguments

import (
	"fmt"
	"regexp"
)

func (arg *Argument) String(flag string, defaultValue string, help string, validationRegEx string) *Argument {

	pattern := regexp.MustCompile(validationRegEx)

	arg.args[flag] = &ArgumentDescriptor[int64, uint64, float64, string, bool]{
		flag:  flag,
		class: String,
		help:  help,
		bounds: PatternBoundary[int64, uint64, float64, string, bool]{
			pattern: pattern,
		},
	}

	if err := arg.args[flag].SetValueString(&defaultValue); err != nil {
		arg.err = fmt.Errorf(defaultValueMustPassValidation, flag)
	}

	return arg

}
