package arguments

func New() *Arguments {
	return &Arguments{
		err:  nil,
		args: make(map[string]ArgumentDescriptor[int64, uint64, float64, string, bool]),
	}
}
