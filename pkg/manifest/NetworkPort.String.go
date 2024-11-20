package manifest

import "fmt"

// String - return the string representation of a network port
func (n *NetworkPort) String() string {
	return fmt.Sprintf("%d", *n)
}
