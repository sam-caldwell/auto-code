package arguments

type ArgumentDescriptor[I int64, U uint64, F float64, S string, B bool] struct {
	flag   string
	class  ArgumentClass
	value  any
	help   string
	bounds Boundary[I, U, F, S, B]
}

type ArgumentClass uint8

const (
	Undefined ArgumentClass = iota
	Bool
	Directory
	Email
	File
	Flag
	Float
	Int
	Secret
	String
	Uint
)
