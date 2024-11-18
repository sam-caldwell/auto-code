package manifest

// GetRef - Returns DocumentReference pointer
func (a *DataObjectWithReference) GetRef() *DocumentReference {
	return &a.Ref
}
