package arguments

func New() *Argument {
	return &Argument{
		err:  nil,
		args: make(map[string]*ArgumentDescriptor[int64, uint64, float64, string, bool]),
	}
}
