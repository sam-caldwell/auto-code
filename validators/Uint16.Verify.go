package validator

// Verify - Verify the Uint16 YAML Manifest Validator Parameter
//
// Things to verify...
// - Make sure min < max
// - verify the parameter is big enough to contain its data range.
func (p Uint16) Verify() error {
	if err := minMaxTypeCheck(p.Min, p.Max, "uint16"); err != nil {
		return err
	}
	return nil
}
