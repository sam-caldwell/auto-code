package validator

// Interface - an abstract interface for many different Interface types (e.g. regex, integer, float)
type Interface interface {
	Valid() bool
	Set(q any) error
}
