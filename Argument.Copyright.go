package arguments

func (arg *Argument) Copyright(copyright string) *Argument {
	arg.copyright = copyright
	return arg
}
