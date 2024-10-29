package arguments

import "fmt"

func (arg *Argument) Int(flag string, defaultValue int64, help string, min int64, max int64) *Argument {

	arg.args[flag] = &ArgumentDescriptor[int64, uint64, float64, string, bool]{
		flag:  flag,
		class: Int,
		value: defaultValue,
		help:  help,
		bounds: NumericBoundary[int64, uint64, float64, string, bool]{
			min: min,
			max: max,
		},
	}
	if err := arg.args[flag].SetValueInt(defaultValue); err != nil {
		arg.err = fmt.Errorf(defaultValueMustPassValidation, flag)
	}
	return arg
}
