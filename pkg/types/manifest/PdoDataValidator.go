package manifest

// PdoDataValidator - an abstract interface for many different validator types (e.g. regex, integer, float)
type PdoDataValidator interface {
	Valid() bool
	Set(q any) error
}
