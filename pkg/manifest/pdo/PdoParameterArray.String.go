package manifest

import "fmt"

// String - return a string representation of a PdoParameterArray
func (p PdoParameterArray) String() (result string) {
	for _, v := range p {
		result += fmt.Sprintf("%v", v)
	}
	return result
}
