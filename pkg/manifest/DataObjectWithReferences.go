package manifest

// DataObjectWithReferences - interface representing all data objects with $ref fields
type DataObjectWithReferences interface {
	GetRef() *DocumentReference
}

// DataObjectWithReference - implements common data object feature with $ref
type DataObjectWithReference struct {
	Ref DocumentReference `yaml:"$ref,omitempty"`
}
