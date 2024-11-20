package artifact

import (
	"fmt"
	"regexp"
	"strings"
)

// Validate - detect and return an error if the Name field is empty or invalid.
func (n *Name) Validate() error {

	const pattern = `^[a-zA-Z][a-zA-Z0-9\-_\.]{0,14}[a-zA-Z0-9]$`

	*n = Name(strings.TrimSpace(string(*n)))

	if re := regexp.MustCompile(pattern); !re.MatchString(string(*n)) {
		return fmt.Errorf("invalid artifact name (%s)", *n)
	}

	return nil
}
