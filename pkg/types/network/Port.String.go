package network

import "fmt"

// String - return the string representation of a network port
func (n *Port) String() string {

	return fmt.Sprintf("%d", *n)

}
