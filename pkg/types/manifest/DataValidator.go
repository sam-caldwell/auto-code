package manifest

// DataValidator - an abstract interface for many different validator types (e.g. regex, integer, float)
type DataValidator interface {
	Valid() bool
	Set(q any)
}
