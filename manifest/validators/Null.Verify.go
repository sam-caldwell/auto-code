package validator

// Verify - Verify the Null (noop) YAML Manifest Validator Parameter
//
// This method will always verify its state.  Everything is always happy.
func (null *Null) Verify() error {
	return nil
}
