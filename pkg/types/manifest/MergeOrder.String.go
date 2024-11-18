package manifest

// String - return string representation of MergeOrder
func (m *MergeOrder) String() string {
	return [...]string{"files", "environment", "command-line"}[*m]
}
