package arguments

// StringDataObject - represents a string value and its validator as an implementation of ValueDataInterface.
type StringDataObject struct {
	data      any
	validator *Validator
	err       error
}

// Action - validates the given value using the associated validator.
//
// This method calls the validator's action on the provided value and returns
// the result. If no validator is set, it returns false by default.
func (v StringDataObject) Action(value any) bool {
	if v.validator == nil {
		return false // No validator, so assume invalid.
	}
	return v.validator.action(value) // Hypothetical method call on validator
}

// Data - returns the stored data value.
//
// This method returns the data stored in the ValueDataObject instance.
func (v StringDataObject) Data() any {
	return v.data
}

// Error - return the object error state
func (v StringDataObject) Error() error {
	return v.err
}
