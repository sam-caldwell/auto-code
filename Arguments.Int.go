package arguments

func (arg *Arguments) Int(flag string, defaultValue int64, help string, min int64, max int64) *Arguments {
	arg.args = append(
		arg.args,
		ArgumentDescriptor[int64, uint64, float64, string, bool]{
			flag:  flag,
			value: defaultValue,
			help:  help,
			bounds: NumericBoundary[int64, uint64, float64, string, bool]{
				min: min,
				max: max,
			},
		})
	return arg
}
