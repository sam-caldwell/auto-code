package validator

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"regexp"
	"strings"
)

// Verify - Verify the Pattern
func (p *Pattern) Verify() error {
	if strings.TrimSpace(p.Regex) == "" {
		return fmt.Errorf(messages.ErrEmptyRegularExpression)
	}
	_, err := regexp.Compile(string(p.Regex))

	return err

}
