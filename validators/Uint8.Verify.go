package validator

// Verify - Verify the Uint8 YAML Manifest Validator Parameter
//
// Things to verify...
// - Make sure min < max
// - verify the parameter is big enough to contain its data range.
func (p Uint8) Verify() error {
	if err := minMaxTypeCheck(p.Min, p.Max, "uint8"); err != nil {
		return err
	}
	return nil
}
