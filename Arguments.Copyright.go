package arguments

func (arg *Arguments) Copyright(copyright string) *Arguments {
	arg.copyright = copyright
	return arg
}
