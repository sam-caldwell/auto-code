package pdo

import "fmt"

// String - return a string representation of a Array
func (p Array) String() (result string) {
	for _, v := range p {
		result += fmt.Sprintf("%v", v)
	}
	return result
}
