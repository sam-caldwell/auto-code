package httpDriver

// HttpDriver - an enumerated type describing the http endpoint driver
type HttpDriver uint8

const (
	Rest HttpDriver = iota
	GraphQl
)
