package generic

import (
	"fmt"
	"regexp"
	"strings"
)

// Validate - detect and return an error if the Name field is empty or invalid.
func (p *Identifier) Validate() error {

	const pattern = `^[a-zA-Z][a-zA-Z0-9\-_\.]{0,14}[a-zA-Z0-9]$`

	*p = Identifier(strings.TrimSpace(string(*p)))

	if re := regexp.MustCompile(pattern); !re.MatchString(string(*p)) {
		return fmt.Errorf("invalid artifact name (%s)", *p)
	}

	return nil
}
