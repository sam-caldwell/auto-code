package arguments

type ArgumentDescriptor[I int64, U uint64, F float64, S string, B bool] struct {
	flag   string
	class  ArgumentClass
	value  any
	help   string
	bounds Boundary[I, U, F, S, B]
}
