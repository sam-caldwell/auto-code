package manifest

import "strings"

// String - return a string representation of an enumerated type
func (p PdoParameterEnum) String() string {
	return strings.Join(p, ",")
}
