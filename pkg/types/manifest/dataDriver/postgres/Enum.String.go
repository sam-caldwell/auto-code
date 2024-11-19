package postgres

import "strings"

// String - return the string representation of an enumerated type (value set)
func (e *Enum) String() string {
	return strings.Join(*e, ",")
}
