package validator

// Verify - Verify the Uint64 YAML Manifest Validator Parameter
//
// Things to verify...
// - Make sure min < max
// - verify the parameter is big enough to contain its data range.
func (p Uint64) Verify() error {
	if err := minMaxTypeCheck(p.Min, p.Max, "uint64"); err != nil {
		return err
	}
	return nil
}
