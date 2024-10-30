package arguments

type InvalidDataObject struct {
	err error
}

// Action - noop validator action
func (v InvalidDataObject) Action(value any) bool {
	return false
}

// Data - noop data getter
func (v InvalidDataObject) Data() any {
	return nil
}

// Error - return error state
func (v InvalidDataObject) Error() error {
	return v.err
}
