package xref

// DataObjectWithReferences - interface representing all dataContract objects with $ref fields
type DataObjectWithReferences interface {
	GetRef() *DocumentReference
}

// DataObjectWithReference - implements common dataContract object feature with $ref
type DataObjectWithReference struct {
	Ref DocumentReference `yaml:"$ref,omitempty"`
}
