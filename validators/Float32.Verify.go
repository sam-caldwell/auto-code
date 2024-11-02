package validator

// Verify - Verify the Int YAML Manifest Validator Parameter
//
// Things to verify...
// - Make sure min < max
// - verify the parameter is big enough to contain its data range.
func (p Float32) Verify() error {
	if err := minMaxTypeCheck(p.Min, p.Max, "float32"); err != nil {
		return err
	}
	return nil
}
