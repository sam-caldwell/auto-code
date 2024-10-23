package arguments

type Arguments struct {
	err         error
	programName string
	copyright   string
	args        []ArgumentDescriptor[int64, uint64, float64, string, bool]
}
