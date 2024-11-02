package validator

// Verify - Verify the Int32 YAML Manifest Validator Parameter
//
// Things to verify...
// - Make sure min < max
// - verify the parameter is big enough to contain its data range.
func (p Int32) Verify() error {
	if err := minMaxTypeCheck(p.Min, p.Max, "int32"); err != nil {
		return err
	}
	return nil
}
