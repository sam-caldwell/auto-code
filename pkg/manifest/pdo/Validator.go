package pdo

// Validator - an abstract interface for many different validator types (e.g. regex, integer, float)
type Validator interface {
	Valid() bool
	Set(q any) error
}
