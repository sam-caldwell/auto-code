package arguments

func (arg *Argument) ProgramName(name string) *Argument {
	arg.programName = name
	return arg
}
