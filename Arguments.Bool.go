package arguments

func (arg *Arguments) Bool(flag string, defaultValue bool, help string) *Arguments {
	arg.args = append(
		arg.args,
		ArgumentDescriptor[int64, uint64, float64, string, bool]{
			flag:   flag,
			value:  defaultValue,
			help:   help,
			bounds: NumericBoundary[int64, uint64, float64, string, bool]{},
		})
	return arg
}
