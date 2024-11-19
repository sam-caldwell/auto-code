package manifest

// String - return a string representation of HttpEndpointMethod
func (h *HttpEndpointMethod) String() string {
	switch *h {
	case GET:
		return "get"
	case POST:
		return "post"
	case PUT:
		return "put"
	case DELETE:
		return "delete"
	default:
		panic("unknown or unsupported Http Method")
	}
}
