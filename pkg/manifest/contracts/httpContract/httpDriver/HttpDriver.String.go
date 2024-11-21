package httpDriver

// String - return a string representation of an HttpDriver
func (e *HttpDriver) String() string {
	switch *e {
	case Rest:
		return "rest"
	case GraphQl:
		return "graphql"
	default:
		panic("invalid http driver")
	}
}
