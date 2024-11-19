package manifest

import (
	"fmt"
	"regexp"
)

// Set - define the regular expression input
func (p *PdoValidatorRegex) Set(input any) error {
	switch input.(type) {
	case *regexp.Regexp:
		p.re = input.(*regexp.Regexp)
	default:
		return fmt.Errorf("input type not supported (expect regular expression)")
	}
	return nil
}
