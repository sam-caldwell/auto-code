package arguments

func (arg *Arguments) ProgramName(name string) *Arguments {
	arg.programName = name
	return arg
}
