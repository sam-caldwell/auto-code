package arguments

type ArgumentDescriptor[I int64, U uint64, F float64, S String, B bool] struct {
	flag   string
	value  any
	help   string
	bounds Boundary[I, U, F, S, B]
}
