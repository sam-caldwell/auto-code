package arguments

type Boundary[I int64, U uint64, F float64, S string, B bool] interface {
	Validator(value any) bool
}
