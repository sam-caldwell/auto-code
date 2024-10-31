package arguments

// dsInvalidParameter - Represents an error state when an invalid data source specification is encountered.
//
// This signals to the caller of any DataSource that the data source was improperly configured and
// should not be trusted.  It is invalid.
type dsInvalidParameter struct {
	err error
}

// Get - This will return the error state and a nil data object.
//
// In reality, we should never get to the point where we would use this.
// This exists to make the interface happy.
func (ds dsInvalidParameter) Get() (any, error) {
	return nil, ds.err
}

// Error - Return the object's error state
func (ds dsInvalidParameter) Error() error {
	return ds.err
}
