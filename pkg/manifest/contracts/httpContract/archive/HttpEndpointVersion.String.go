package httpContract

import "fmt"

// String - return a string representation of the HttpEndpointVersion using the 'v${N}' format.
func (h *HttpEndpointVersion) String() string {
	return fmt.Sprintf("v%d", *h)
}
