package manifest

import (
	"fmt"
	"strings"
)

// String - return a string representation of an enumerated type
//
// An enum is represented as follows: 'enum(v1,v2,v3)'
func (p PdoParameterEnum) String() string {
	return fmt.Sprintf("enum(%s)", strings.Join(p, ","))
}
