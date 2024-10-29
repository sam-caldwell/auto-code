package arguments

func (arg *Argument) Flag(flag string, help string) *Argument {

	arg.args[flag] = &ArgumentDescriptor[int64, uint64, float64, string, bool]{
		flag:   flag,
		class:  Flag,
		value:  false,
		help:   help,
		bounds: NoBoundaries[int64, uint64, float64, string, bool]{},
	}

	return arg

}
