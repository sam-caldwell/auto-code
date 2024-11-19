package postgres

import (
	"fmt"
	"strings"
)

// String - return the string representation of an enumerated type (value set)
func (e *Enum) String() string {
	return fmt.Sprintf("%s(%s)", e.Name, strings.Join(e.Elements, ","))
}
