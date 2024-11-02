package validator

// Verify - Verify the Uint32 YAML Manifest Validator Parameter
//
// Things to verify...
// - Make sure min < max
// - verify the parameter is big enough to contain its data range.
func (p Uint32) Verify() error {
	if err := minMaxTypeCheck(p.Min, p.Max, "uint32"); err != nil {
		return err
	}
	return nil
}
