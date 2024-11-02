package validator

import (
	"fmt"
	"regexp"
	"strings"
)

// Verify - Verify the Pattern
func (p *Pattern) Verify() error {
	if strings.TrimSpace(p.Regex) == "" {
		return fmt.Errorf("invalid (empty) regular expression")
	}
	_, err := regexp.Compile(string(p.Regex))

	return err

}
