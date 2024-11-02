package validator

import "regexp"

// Verify - Verify the Pattern
func (p *Pattern) Verify() error {

	_, err := regexp.Compile(string(p.Regex))

	return err

}
