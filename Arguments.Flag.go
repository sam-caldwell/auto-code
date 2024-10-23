package arguments

func (arg *Arguments) Flag(flag string, help string) *Arguments {

	arg.args[flag] = ArgumentDescriptor[int64, uint64, float64, string, bool]{
		flag:   flag,
		class:  Flag,
		value:  false,
		help:   help,
		bounds: NoBoundaries[int64, uint64, float64, string, bool]{},
	}

	return arg

}
