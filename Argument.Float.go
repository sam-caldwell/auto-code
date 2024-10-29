package arguments

func (arg *Argument) Float(flag string, defaultValue float64, help string, min, max float64) *Argument {

	arg.args[flag] = &ArgumentDescriptor[int64, uint64, float64, string, bool]{
		flag:  flag,
		value: defaultValue,
		help:  help,
		bounds: NumericBoundary[int64, uint64, float64, string, bool]{
			min: min,
			max: max,
		},
	}
	return arg
}
