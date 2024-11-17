package manifest

// String - return the string representation of the SemanticVersion
func (sv *SemanticVersion) String() string {
	return string(*sv)
}
