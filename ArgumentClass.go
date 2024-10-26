package arguments

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
