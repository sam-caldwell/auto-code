package manifest

import (
	"fmt"
	"strings"
)

// String - return a string representation of an enumerated type
func (p PdoParameterEnum) String() string {
	return fmt.Sprintf("enum(%s)", strings.Join(p, ","))
}
