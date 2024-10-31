package arguments

import "strings"

func (arg *Argument) Bool(flag string, defaultValue bool, help string) *Argument {

	name := strings.TrimPrefix("--", strings.TrimPrefix(flag, "-"))
	arg.args[name] = &ArgumentDescriptor[int64, uint64, float64, string, bool]{
		flag:   flag,
		class:  Bool,
		value:  defaultValue,
		help:   help,
		bounds: NumericBoundary[int64, uint64, float64, string, bool]{},
	}
	return arg
}
